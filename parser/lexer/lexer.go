package lexer

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/borschtapp/kapusta/dictionary"
)

const eof = -1

type Lexer struct {
	input             string              // the string being scanned
	dict              *dictionary.Dict    // the language of the string
	start             int                 // start position of this Token
	pos               int                 // current position of the input
	width             int                 // width of last rune read from input
	prev              TokenType           // previous token type
	items             chan Token          // channel of scanned items
	done              chan struct{}       // closed by Close() to unblock the goroutine
	ingredientMatcher *dictionary.Matcher // optional matcher for known ingredient names
}

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*Lexer) stateFn

// Lex creates a new Lexer
func Lex(input string, lang string) (*Lexer, error) {
	return LexWithOptions(input, lang, nil)
}

// LexWithOptions creates a new Lexer with an optional list of known ingredient names.
// The lexer will emit ItemIngredient tokens when any of them are matched in the input.
func LexWithOptions(input string, lang string, knownIngredients []string) (*Lexer, error) {
	dict, err := dictionary.ForLang(lang)
	if err != nil {
		return nil, err
	}
	var matcher *dictionary.Matcher
	if len(knownIngredients) > 0 {
		matcher = dictionary.NewMatcher(knownIngredients)
	}
	l := &Lexer{
		input:             input,
		dict:              dict,
		items:             make(chan Token),
		done:              make(chan struct{}),
		ingredientMatcher: matcher,
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
		Type:       t,
		Lexeme:     l.input[l.start:l.pos],
		Code:       code,
		Value:      val,
		StartIndex: l.start,
		EndIndex:   l.pos,
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
		Type:       ItemError,
		Lexeme:     fmt.Sprintf(format, args...),
		StartIndex: l.start,
		EndIndex:   l.pos,
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
