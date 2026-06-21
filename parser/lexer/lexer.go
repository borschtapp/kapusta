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
	tokens            []Token             // slice of scanned items
	cursor            int                 // current position in tokens slice
	ingredientMatcher *dictionary.Matcher // optional matcher for known ingredient names
	score             int                 // total confidence score based on recognized tokens
}

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*Lexer) stateFn

// Lex creates a new Lexer
func Lex(input string, lang string) (*Lexer, error) {
	return LexWithMatcher(input, lang, nil)
}

// LexWithMatcher creates a new Lexer with an optional pre-built matcher for known ingredient names.
// The lexer will emit ItemIngredient tokens when any of them are matched in the input.
func LexWithMatcher(input string, lang string, matcher *dictionary.Matcher) (*Lexer, error) {
	dict, err := dictionary.ForLanguage(lang)
	if err != nil {
		return nil, err
	}
	l := &Lexer{
		input:             input,
		dict:              dict,
		tokens:            make([]Token, 0, 16),
		ingredientMatcher: matcher,
	}
	l.run()
	return l, nil
}

// Next returns the next Token from the input.
func (l *Lexer) Next() (Token, error) {
	if l.cursor >= len(l.tokens) {
		return Token{Type: ItemEOF}, nil
	}
	tok := l.tokens[l.cursor]
	l.cursor++
	if tok.Type == ItemError {
		return tok, errors.New(tok.Lexeme)
	}
	return tok, nil
}

// run runs the lexer.
func (l *Lexer) run() {
	for state := lexInsideAction; state != nil; {
		state = state(l)
	}
}

func (l *Lexer) emit(t TokenType) {
	l.emitToken(Token{
		Type: t,
	})
}

func (l *Lexer) emitValue(t TokenType, code string, val float64) {
	l.emitToken(Token{
		Type:  t,
		Code:  code,
		Value: val,
	})
}

func (l *Lexer) emitToken(tok Token) {
	tok.StartIndex = l.start
	tok.EndIndex = l.pos
	tok.Lexeme = l.input[l.start:l.pos]

	l.score += len(tok.Lexeme) * tok.Type.Weight()

	l.tokens = append(l.tokens, tok)
	l.start = l.pos
	l.prev = tok.Type
}

// Score returns the total confidence score calculated during lexing.
func (l *Lexer) Score() int {
	return l.score
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
	l.tokens = append(l.tokens, Token{
		Type:       ItemError,
		Lexeme:     fmt.Sprintf(format, args...),
		StartIndex: l.start,
		EndIndex:   l.pos,
	})
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
