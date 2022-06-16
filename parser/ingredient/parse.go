package ingredient

import (
	"fmt"

	"github.com/borschtapp/krip/model"
	"github.com/borschtapp/krip/utils"
)

func Parse(str string, lang string) (*model.Ingredient, error) {
	l := Lex(str, lang)

	var unit, text string
	var prevTokType tokenType
	ingredient := model.Ingredient{}
	for tok, err, eof := l.Next(); !eof; tok, err, eof = l.Next() {
		if err != nil {
			return nil, fmt.Errorf("failed to parse ingredient: %v", err)
		}

		if tok.Type == itemNumber {
			if ingredient.Quantity == 0 {
				if val, err := utils.ParseFloat(tok.Lexeme); err == nil {
					ingredient.Quantity = val
				} else {
					return nil, fmt.Errorf("failed to parse ingredient amount: %v", err)
				}
			} else if prevTokType == itemIdentifierRange {
				if val, err := utils.ParseFloat(tok.Lexeme); err == nil {
					ingredient.QuantityMax = val
				} else {
					return nil, fmt.Errorf("failed to parse ingredient amount: %v", err)
				}
			} else {
				return nil, fmt.Errorf("amount already set in source [%s]", str)
			}
		} else if tok.Type == itemNumberFraction {
			if ingredient.Quantity == 0 || prevTokType == itemNumber {
				if val, err := utils.ParseFraction(tok.Lexeme); err == nil {
					ingredient.Quantity += val
				} else {
					return nil, fmt.Errorf("failed to parse ingredient amount (fr): %v", err)
				}
			} else if prevTokType == itemIdentifierRange {
				if val, err := utils.ParseFraction(tok.Lexeme); err == nil {
					ingredient.QuantityMax = val
				} else {
					return nil, fmt.Errorf("failed to parse ingredient amount: %v", err)
				}
			} else {
				return nil, fmt.Errorf("amount already set (fr) in source [%s]", str)
			}
		} else if tok.Type == itemUnit {
			if text != "" {
				unit += text
				text = ""
			}
			if unit != "" {
				unit += " "
			}
			unit += tok.Lexeme
		} else if tok.Type == itemIdentifier {
			if text != "" {
				text += " "
			}
			text += tok.Lexeme
		} else if tok.Type == itemComment {
			ingredient.Annotation = tok.Lexeme
		}

		prevTokType = tok.Type
	}
	ingredient.Unit = unit
	ingredient.Ingredient = text
	return &ingredient, nil
}
