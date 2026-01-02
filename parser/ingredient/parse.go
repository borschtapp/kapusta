package ingredient

import (
	"fmt"
	"strings"

	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/krip/utils"
)

func Parse(str string, lang string) (*model.Ingredient, error) {
	// replace some crazy characters
	str = strings.ReplaceAll(str, "⁄", "/")

	l := Lex(str, lang)

	var unit, text string
	secondQuantity := false // as of now, we ignore second quantity
	var prevTokType tokenType
	ingredient := model.Ingredient{}
	for tok, err, eof := l.Next(); !eof; tok, err, eof = l.Next() {
		if err != nil {
			return nil, fmt.Errorf("failed to parse ingredient: %v", err)
		}

		switch tok.Type {
		case itemNumber:
			if ingredient.Quantity == 0 {
				if val, err := utils.ParseFloat(tok.Lexeme); err == nil {
					ingredient.Quantity = val
				} else {
					tok.Type = itemIdentifier
				}
			} else if prevTokType == itemIdentifierRange {
				if val, err := utils.ParseFloat(tok.Lexeme); err == nil {
					ingredient.QuantityMax = val
				} else {
					return nil, fmt.Errorf("failed to parse ingredient amount: %v", err)
				}
			} else if prevTokType == itemSep {
				secondQuantity = true
			} else {
				tok.Type = itemIdentifierSkip
			}
		case itemNumberFraction:
			if ingredient.Quantity == 0 || prevTokType == itemNumber {
				if val, err := utils.ParseFraction(tok.Lexeme); err == nil {
					ingredient.Quantity += val
				} else {
					tok.Type = itemIdentifier
				}
			} else if prevTokType == itemIdentifierRange {
				if val, err := utils.ParseFraction(tok.Lexeme); err == nil {
					ingredient.QuantityMax = val
				} else {
					return nil, fmt.Errorf("failed to parse ingredient amount: %v", err)
				}
			} else if prevTokType == itemSep {
				secondQuantity = true
			} else {
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
		case tok.Type == itemComment:
			ingredient.Annotation = tok.Lexeme
		}

		prevTokType = tok.Type
	}

	// split text if it contains comma
	if strings.Contains(text, ",") && ingredient.Annotation == "" {
		split := strings.SplitN(text, ",", 2)
		text = strings.TrimSpace(split[0])
		ingredient.Annotation = strings.TrimSpace(split[1])
	}

	ingredient.Unit = unit
	ingredient.Ingredient = text
	return &ingredient, nil
}
