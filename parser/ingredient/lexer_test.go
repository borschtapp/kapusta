package ingredient

import (
	"testing"
)

func TestLex(t *testing.T) {
	l := Lex("2 cups warm water divided", "en")

	for tok, err, eof := l.Next(); !eof; tok, err, eof = l.Next() {
		if err != nil {
			t.Errorf("failed to parse ingredient: %v", err)
			return
		}
		_ = tok
	}
}

func TestLexUnknownLanguage(t *testing.T) {
	l := Lex("2 cups warm water", "non-existent-lang")
	foundError := false
	for tok, _, eof := l.Next(); !eof; tok, _, eof = l.Next() {
		if tok.Type == itemError {
			foundError = true
		}
	}
	if !foundError {
		t.Error("Expected an error token for unknown language, but got none")
	}
}
