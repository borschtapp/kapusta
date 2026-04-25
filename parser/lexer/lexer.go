package lexer

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/borschtapp/kapusta/dictionary"
)

type Token struct {
	Type        TokenType
	Lexeme      string
	Code        string
	Value       float64
	StartColumn int
	EndColumn   int
}

func (i Token) String() string {
	if i.Type == ItemEOF {
		return "EOF"
	} else if i.Type == ItemError {
		return i.Lexeme
	} else if len(i.Lexeme) > 10 {
		return fmt.Sprintf("%.10q...", i.Lexeme)
	}
	return fmt.Sprintf("%q", i.Lexeme)
}

type TokenType int

func (it TokenType) IsNumber() bool {
	return it == ItemNumber || it == ItemNumberFraction
}

func (it TokenType) String() string {
	switch it {
	case ItemIdentifier:
		return "IDENT"
	case ItemComment:
		return "COMMENT"
	case ItemNumber:
		return "NUMBER"
	case ItemNumberFraction:
		return "FRACTION"
	case ItemSep:
		return "SEPARATOR"
	case ItemUnit:
		return "UNIT"
	case ItemSizeSuffix:
		return "SIZE_SUFFIX"
	default:
		return fmt.Sprintf("Unknown [%d]", it)
	}
}

// TokenType identifies the type of lex items
const (
	ItemError TokenType = iota // error occurred; value is text of error
	ItemEOF
	ItemIdentifier
	ItemComment
	ItemNumber
	ItemNumberFraction
	ItemIdentifierRange
	ItemSep
	ItemUnit
	ItemSizeSuffix
)

const eof = -1

type Lexer struct {
	input string           // the string being scanned
	dict  *dictionary.Dict // the language of the string
	start int              // start position of this Token
	pos   int              // current position of the input
	width int              // width of last rune read from input
	prev  TokenType        // previous token type
	items chan Token       // channel of scanned items
	done  chan struct{}    // closed by Close() to unblock the goroutine
}

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*Lexer) stateFn

// Lex creates a new Lexer
func Lex(input string, lang string) (*Lexer, error) {
	dict, err := dictionary.ForLang(lang)
	if err != nil {
		return nil, err
	}
	l := &Lexer{
		input: input,
		dict:  dict,
		items: make(chan Token),
		done:  make(chan struct{}),
	}
	go l.run()
	return l, nil
}

// Close signals the lexer goroutine to stop if it is blocked sending tokens.
// Callers that do not drain the lexer to EOF must call Close to prevent a goroutine leak.
func (l *Lexer) Close() {
	select {
	case <-l.done:
	default:
		close(l.done)
	}
}

// Next returns the next Token from the input. The Lexer has to be drained
// (all items received until ItemEOF or ItemError) - otherwise the Lexer goroutine will leak.
func (l *Lexer) Next() (Token, error) {
	tok, ok := <-l.items
	if !ok || tok.Type == ItemEOF {
		return Token{Type: ItemEOF}, nil
	}
	if tok.Type == ItemError {
		return tok, errors.New(tok.Lexeme)
	}
	return tok, nil
}

// run runs the lexer - should be run in a separate goroutine.
func (l *Lexer) run() {
	defer close(l.items)
	for state := lexInsideAction; state != nil; {
		state = state(l)
	}
}

func (l *Lexer) emit(t TokenType) {
	l.emitValue(t, "", 0)
}

func (l *Lexer) emitValue(t TokenType, code string, val float64) {
	select {
	case l.items <- Token{
		Type:        t,
		Lexeme:      l.input[l.start:l.pos],
		Code:        code,
		Value:       val,
		StartColumn: l.start,
		EndColumn:   l.pos,
	}:
		l.start = l.pos
		l.prev = t
	case <-l.done:
	}
}

// scan advances to the scan rune in input and returns it
func (l *Lexer) scan() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// ignore skips over the pending input before this point
func (l *Lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune. Can be called only once per call of scan.
func (l *Lexer) backup() {
	l.pos -= l.width
}

// peek returns but does not consume the scan run in the input.
func (l *Lexer) peek() rune {
	r := l.scan()
	l.backup()
	return r
}

// errorf returns an error token and terminates the scan by passing back a nil pointer that will be the next state.
func (l *Lexer) errorf(format string, args ...any) stateFn {
	select {
	case l.items <- Token{
		Type:        ItemError,
		Lexeme:      fmt.Sprintf(format, args...),
		StartColumn: l.start,
		EndColumn:   l.pos,
	}:
	case <-l.done:
	}
	return nil
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.scan()) {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.scan()) {
	}
	l.backup()
}
