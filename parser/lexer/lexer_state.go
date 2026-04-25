package lexer

import (
	"strings"
	"unicode"

	"github.com/borschtapp/kapusta/parser/util"
)

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
			l.emit(ItemEOF)
			return nil
		case unicode.IsSpace(r):
			l.ignore()
		case r == '/' || r == '|':
			l.emit(ItemSep)
		case r == '(':
			return lexBracket
		case r == '–' || r == '—': // en dash / em dash as range separator
			if l.prev.IsNumber() {
				l.emit(ItemIdentifierRange)
			} else {
				l.ignore()
			}
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
		case util.IsFraction(r):
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
	if l.accept("/") || l.accept(util.Fractions) {
		return lexFractions
	}
	if l.peek() == '-' {
		l.scan()
		isRange := unicode.IsDigit(l.peek())
		l.backup()
		if isRange {
			val, _ := util.ParseFloat(l.input[l.start:l.pos])
			l.emitValue(ItemNumber, "", val)
			l.accept("-")
			l.emit(ItemIdentifierRange)
			return lexInsideAction
		}
	}

	val, _ := util.ParseFloat(l.input[l.start:l.pos])
	l.emitValue(ItemNumber, "", val)
	return lexInsideAction
}

func lexFractions(l *Lexer) stateFn {
	digits := "0123456789"
	l.acceptRun(digits)
	l.accept(util.Fractions)

	if l.accept("/") {
		l.acceptRun(digits)
	}

	val, _ := util.ParseFraction(l.input[l.start:l.pos])
	l.emitValue(ItemNumberFraction, "", val)
	return lexInsideAction
}

// lexBracket scans a parenthesised comment, handling arbitrary nesting.
func lexBracket(l *Lexer) stateFn {
	depth := 1
	for depth > 0 {
		switch l.scan() {
		case eof:
			return l.errorf("unterminated string in bracket")
		case '(':
			depth++
		case ')':
			depth--
		}
	}

	l.emitValue(ItemComment, "", 0)
	return lexInsideAction
}

// StripParens recursively removes matching outer parentheses and trims whitespace.
func StripParens(s string) string {
	for {
		s = strings.TrimSpace(s)
		if len(s) < 2 || s[0] != '(' || s[len(s)-1] != ')' {
			return s
		}

		// Verify the outer parentheses are a matching pair
		depth := 0
		for i, r := range s {
			if r == '(' {
				depth++
			} else if r == ')' {
				depth--
			}
			// If depth reaches zero before the end of the string, the outer parens aren't a pair
			if depth == 0 && i < len(s)-1 {
				return s
			}
		}
		s = s[1 : len(s)-1]
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
		l.emitValue(ItemUnit, code, 0)
		return lexInsideAction
	}

	ident := l.input[l.start:l.pos]
	lower := strings.ToLower(ident)
	if _, ok := l.dict.FindSizeSuffix(lower); ok {
		l.emit(ItemSizeSuffix)
		return lexInsideAction
	}

	if val, ok := l.dict.FindNumber(lower); ok {
		l.emitValue(ItemNumber, "", val)
	} else if _, ok := l.dict.FindQuantityBetween(lower); l.prev.IsNumber() && ok {
		l.emit(ItemIdentifierRange)
	} else {
		l.emit(ItemIdentifier)
	}
	return lexInsideAction
}
