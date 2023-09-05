package ingredient

import (
	"errors"
	"fmt"
	"github.com/borschtapp/kapusta/dictionary"
	"strings"
	"unicode/utf8"
)

type Token struct {
	Type        tokenType
	Lexeme      string
	StartColumn int
	EndColumn   int
}

func (i Token) String() string {
	switch i.Type {
	case itemEOF:
		return "EOF"
	case itemError:
		return i.Lexeme
	}
	if len(i.Lexeme) > 10 {
		return fmt.Sprintf("%.10q...", i.Lexeme)
	}
	return fmt.Sprintf("%q", i.Lexeme)
}

type tokenType int

func (it tokenType) String() string {
	switch it {
	case itemIdentifier:
		return "IDENT"
	case itemComment:
		return "COMMENT"
	case itemNumber:
		return "NUMBER"
	case itemNumberFraction:
		return "FRACTION"
	case itemSep:
		return "SEPARATOR"
	case itemUnit:
		return "UNIT"
	default:
		return fmt.Sprintf("Unknown [%d]", it)
	}
}

// tokenType identifies the type of lex items
const (
	itemError tokenType = iota // error occurred; value is text of error
	itemEOF
	itemIdentifier
	itemIdentifierSkip
	itemComment
	itemNumber
	itemNumberFraction
	itemIdentifierRange
	itemSep
	itemUnit
)

const eof = -1

type Lexer struct {
	input string           // the string being scanned
	dict  *dictionary.Dict // the language of the string
	start int              // start position of this Token
	pos   int              // current position of the input
	width int              // width of last rune read from input
	prev  tokenType        // previous token type
	items chan Token       // channel of scanned items
}

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*Lexer) stateFn

// Lex creates a new Lexer
func Lex(input string, lang string) *Lexer {
	dict, err := dictionary.ForLang(lang)

	l := &Lexer{
		input: input,
		dict:  dict,
		items: make(chan Token),
	}
	go l.run(err)
	return l
}

// Next returns the next Token from the input. The Lexer has to be drained
// (all items received until itemEOF or itemError) - otherwise the Lexer goroutine will leak.
func (l *Lexer) Next() (Token, error, bool) {
	tok := <-l.items

	if tok.Type == itemEOF {
		return tok, nil, true
	} else if tok.Type == itemError {
		return tok, errors.New(tok.Lexeme), false
	}
	return tok, nil, false
}

// run runs the lexer - should be run in a separate goroutine.
func (l *Lexer) run(err error) {
	if err != nil {
		l.errorf("%v", err)
	}
	for state := lexInsideAction; state != nil; {
		state = state(l)
	}
	close(l.items) // no more tokens will be delivered
}

func (l *Lexer) emit(t tokenType) {
	l.items <- Token{t, l.input[l.start:l.pos], l.start, l.pos}
	l.start = l.pos
	l.prev = t
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
func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- Token{itemError, fmt.Sprintf(format, args...), l.start, l.pos}
	return nil
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.scan()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.scan()) >= 0 {
	}
	l.backup()
}
