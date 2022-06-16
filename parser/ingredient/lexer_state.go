package ingredient

import (
	"strings"
	"unicode"

	"github.com/borschtapp/krip/utils"
)

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r' || r == '\n'
}

// isFraction reports whether r is a fraction number.
func isFraction(r rune) bool {
	return strings.IndexAny(string(r), utils.Fractions) >= 0
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || r == '-' || r == ',' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

func lexInsideAction(l *Lexer) stateFn {
	// Either number, quoted string or identifier.
	// Spaces separate and are ignored.
	// Pipe symbols separate and are emitted.
	for {
		switch r := l.scan(); {
		case r == eof || r == '\n':
			l.emit(itemEOF)
		case isSpace(r):
			l.ignore()
		case r == '|':
			l.emit(itemPipe)
		case r == '"':
			return lexQuote
		case r == '`':
			return lexRawQuote
		case r == '(':
			return lexBracket
		case r == '+' || r == '-' || '0' <= r && r <= '9':
			l.backup()
			return lexNumber
		case isFraction(r):
			l.backup()
			return lexFractions
		case isAlphaNumeric(r):
			l.backup()
			return lexIdentifier
		default:
			return l.errorf("unrecognized character in action: %#U", r)
		}
	}
}

func lexNumber(l *Lexer) stateFn {
	// Optional leading sign.
	l.accept("+-")
	// is it hex?
	digits := "0123456789"
	if l.accept("0") && l.accept("xX") {
		digits = "0123456789abcdefABCDEF"
	}
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}
	if l.accept("eE") {
		l.accept("+-")
		l.acceptRun("0123456789")
	}
	if l.accept("/") || l.accept(utils.Fractions) {
		return lexFractions
	}
	if l.accept("-") { // range, like :from-:to
		l.backup()
		l.emit(itemNumber)
		l.scan()
		l.emit(itemIdentifierRange)
		return lexInsideAction
	}
	if isAlphaNumeric(l.peek()) {
		l.scan()
		return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
	}
	l.emit(itemNumber)
	return lexInsideAction
}

func lexFractions(l *Lexer) stateFn {
	digits := "0123456789"
	l.acceptRun(digits)
	l.accept(utils.Fractions)

	if l.accept("/") {
		l.acceptRun(digits)
	}

	if isAlphaNumeric(l.peek()) {
		l.scan()
		return l.errorf("bad fraction syntax: %q", l.input[l.start:l.pos])
	}
	l.emit(itemNumberFraction)
	return lexInsideAction
}

func lexQuote(l *Lexer) stateFn {
	for {
		switch l.scan() {
		case '\\':
			if r := l.scan(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated quoted string")
		case '"':
			l.emit(itemString)
			return lexInsideAction
		}
	}
}

// lexRawQuote scans a raw quoted string.
func lexRawQuote(l *Lexer) stateFn {
	for {
		switch l.scan() {
		case eof:
			return l.errorf("unterminated raw quoted string")
		case '`':
			l.emit(itemRawString)
			return lexInsideAction
		}
	}
}

// lexBracket scans a string in bracket.
func lexBracket(l *Lexer) stateFn {
	numBrackets := 1

	l.ignore()
	for {
		switch l.scan() {
		case eof:
			return l.errorf("unterminated string in bracket")
		case '(':
			numBrackets++
			l.ignore()
		case ')':
			l.backup()
			l.emit(itemComment)
			for l.scan() == ')' {
				numBrackets--
				if numBrackets == 0 {
					break
				}
			}
			return lexInsideAction
		}
	}
}

// lexIdentifier scans an alphanumeric.
func lexIdentifier(l *Lexer) stateFn {
	// Advance as long as the current rune is alphanumeric
	for isAlphaNumeric(l.scan()) {
	}

	// Now the current rune is no longer alphanumeric,
	// so back up to the last alphanumeric rune and emit the current item.
	l.backup()

	ident := l.input[l.start:l.pos]
	if _, ok := l.dict.FindUnit(ident); ok {
		l.emit(itemUnit)
	} else if _, ok := l.dict.FindQuantityBetween(ident); l.prev == itemNumber && ok {
		l.emit(itemIdentifierRange)
	} else {
		l.emit(itemIdentifier)
	}
	return lexInsideAction
}
