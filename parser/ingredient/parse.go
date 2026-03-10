package ingredient

import (
	"fmt"
	"strings"

	"github.com/borschtapp/kapusta/model"
)

func Parse(str string, lang string) (*model.Ingredient, error) {
	// replace some crazy characters
	str = strings.ReplaceAll(str, "⁄", "/")

	l := Lex(str, lang)

	var unit, unitCode, text string
	secondQuantity := false // as of now, we ignore second quantity
	var prevTokType tokenType
	ingredient := model.Ingredient{}
	for tok, err, eof := l.Next(); !eof; tok, err, eof = l.Next() {
		if err != nil {
			return nil, fmt.Errorf("failed to parse ingredient: %v", err)
		}

		switch tok.Type {
		case itemNumber, itemNumberFraction:
			if ingredient.Amount == 0 {
				// First number: primary amount
				ingredient.Amount = tok.Value
			} else if tok.Type == itemNumberFraction && prevTokType == itemNumber {
				// Mixed number: e.g. "5 ½" → 5 + 0.5
				ingredient.Amount += tok.Value
			} else if prevTokType == itemIdentifierRange {
				// Range upper bound: e.g. "1-2 cups"
				ingredient.MaxAmount = tok.Value
			} else if prevTokType == itemSep {
				// Second measurement after separator: e.g. "2 cups / 480 ml" — ignore
				secondQuantity = true
			} else {
				// Extra number with no clear role — treat as text
				tok.Type = itemIdentifierSkip
			}
		}

		switch {
		case tok.Type == itemIdentifier || tok.Type == itemIdentifierSkip || (tok.Type == itemUnit && text != ""):
			if text != "" && prevTokType != itemIdentifierSkip {
				text += " "
			}
			text += tok.Lexeme
		case tok.Type == itemUnit && !secondQuantity:
			if unit != "" {
				unit += " "
			}
			unit += tok.Lexeme
			unitCode = tok.Code
		case tok.Type == itemComment:
			ingredient.Description = tok.Lexeme
		}

		prevTokType = tok.Type
	}

	// split text if it contains comma
	if strings.Contains(text, ",") && ingredient.Description == "" {
		split := strings.SplitN(text, ",", 2)
		text = strings.TrimSpace(split[0])
		ingredient.Description = strings.TrimSpace(split[1])
	}

	ingredient.Unit = unit
	ingredient.UnitCode = unitCode
	ingredient.Name = text
	return &ingredient, nil
}
