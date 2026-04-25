package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLex(t *testing.T) {
	l, err := Lex("2 cups warm water divided", "en")
	assert.NoError(t, err)
	for {
		tok, err := l.Next()
		assert.NoError(t, err, "failed to parse ingredient")
		if tok.Type == ItemEOF {
			break
		}
	}
}

func TestLexUnknownLanguage(t *testing.T) {
	_, err := Lex("2 cups warm water", "non-existent-lang")
	assert.Error(t, err, "Expected an error for unknown language, but got none")
}

func TestLexEnDashRange(t *testing.T) {
	for _, input := range []string{"1–2 cups water", "1—2 cups water"} {
		t.Run(input, func(t *testing.T) {
			l, err := Lex(input, "en")
			assert.NoError(t, err)
			var types []TokenType
			for {
				tok, _ := l.Next()
				if tok.Type == ItemEOF {
					break
				}
				types = append(types, tok.Type)
			}
			assert.Contains(t, types, ItemIdentifierRange)
		})
	}
}

func BenchmarkLex(b *testing.B) {
	for b.Loop() {
		l, _ := Lex("1 ½ cups diced tomatoes, divided", "en")
		for {
			tok, _ := l.Next()
			if tok.Type == ItemEOF {
				break
			}
		}
	}
}
