package ingredient

import (
	"fmt"
	"testing"
)

func TestLex(t *testing.T) {
	l := Lex("2 cups warm water divided", "en")

	fmt.Println("Type    | Lexeme")
	fmt.Println("--------+------------")
	for tok, err, eof := l.Next(); !eof; tok, err, eof = l.Next() {
		if err != nil {
			t.Errorf("failed to parse ingredient: %v", err)
			return
		}

		fmt.Printf("%-7v | %-10v\n", tok.Type.String(), tok.Lexeme)
	}
}
