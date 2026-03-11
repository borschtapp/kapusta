package ingredient

import (
	"strings"
	"unicode"

	"github.com/borschtapp/krip/utils"
)

// isFraction reports whether r is a fraction number.
func isFraction(r rune) bool {
	return strings.ContainsRune(utils.Fractions, r)
}

// isAlphaNumeric reports whether r is a letter, digit, mark, punctuation, or symbol character.
func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsMark(r) || unicode.IsPunct(r) || unicode.IsSymbol(r)
}

func lexInsideAction(l *Lexer) stateFn {
	// Either number, quoted string or identifier.
	// Spaces separate and are ignored.
	// Pipe symbols separate and are emitted.
	for {
		switch r := l.scan(); {
		case r == eof || r == '\n':
			l.emit(itemEOF)
		case unicode.IsSpace(r):
			l.ignore()
		case r == '/' || r == '|':
			l.emit(itemSep)
		case r == '(':
			return lexBracket
		case r == '+' || r == '-':
			n := l.peek()
			if (n >= '0' && n <= '9') || n == '.' {
				l.backup()
				return lexNumber
			}
			l.backup()
			return lexIdentifier
		case '0' <= r && r <= '9':
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
	if l.accept("/") || l.accept(utils.Fractions) {
		return lexFractions
	}
	if l.accept("-") { // range, like :from-:to
		l.backup()
		val, _ := utils.ParseFloat(l.input[l.start:l.pos])
		l.emitValue(itemNumber, "", val)
		l.scan()
		l.emit(itemIdentifierRange)
		return lexInsideAction
	}

	val, _ := utils.ParseFloat(l.input[l.start:l.pos])
	l.emitValue(itemNumber, "", val)
	return lexInsideAction
}

func lexFractions(l *Lexer) stateFn {
	digits := "0123456789"
	l.acceptRun(digits)
	l.accept(utils.Fractions)

	if l.accept("/") {
		l.acceptRun(digits)
	}

	val, _ := utils.ParseFraction(l.input[l.start:l.pos])
	l.emitValue(itemNumberFraction, "", val)
	return lexInsideAction
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

// lexIdentifier scans an alphanumeric word and classifies it as a unit, number, range keyword, or plain identifier.
func lexIdentifier(l *Lexer) stateFn {
	// Advance as long as the current rune is alphanumeric
	for isAlphaNumeric(l.scan()) {
	}

	// Now the current rune is no longer alphanumeric,
	// so back up to the last alphanumeric rune and emit the current item.
	l.backup()

	if variant, code, ok := l.dict.FindUnit(l.input[l.start:]); ok {
		l.pos = l.start + len(variant)
		l.emitValue(itemUnit, code, 0)
		return lexInsideAction
	}

	ident := l.input[l.start:l.pos]
	if val, ok := l.dict.FindNumber(ident); ok {
		l.emitValue(itemNumber, "", val)
	} else if _, ok := l.dict.FindQuantityBetween(ident); l.prev == itemNumber && ok {
		l.emit(itemIdentifierRange)
	} else {
		l.emit(itemIdentifier)
	}
	return lexInsideAction
}
