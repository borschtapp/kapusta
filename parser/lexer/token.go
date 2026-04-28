package lexer

import (
	"fmt"
)

// TokenType identifies the type of lex items
const (
	ItemError TokenType = iota // error occurred; value is text of error
	ItemEOF
	ItemIdentifier
	ItemComment
	ItemNumber
	ItemNumberFraction
	ItemIdentifierRange
	ItemSep
	ItemUnit
	ItemSizeSuffix
	ItemTimeUnit
	ItemTemperatureUnit
	ItemIngredient
)

type Token struct {
	Type       TokenType
	Lexeme     string
	Code       string
	Value      float64
	StartIndex int
	EndIndex   int
}

// ShouldMerge returns true if the next token is an adjacent comment that should be
// merged back into this token (e.g. "cup(s)", "гілочку(и)").
func (i Token) ShouldMerge(next Token) bool {
	if next.Type != ItemComment || i.EndIndex != next.StartIndex {
		return false
	}
	switch i.Type {
	case ItemUnit, ItemTimeUnit, ItemTemperatureUnit, ItemIdentifier, ItemIngredient:
		return true
	default:
		return false
	}
}

func (i Token) String() string {
	if i.Type == ItemEOF {
		return "EOF"
	}
	if i.Type == ItemError {
		return i.Lexeme
	}
	if len(i.Lexeme) > 10 {
		return fmt.Sprintf("%.10q...", i.Lexeme)
	}
	return fmt.Sprintf("%q", i.Lexeme)
}

type TokenType int

func (it TokenType) IsNumber() bool {
	return it == ItemNumber || it == ItemNumberFraction
}

func (it TokenType) String() string {
	switch it {
	case ItemIdentifier:
		return "IDENT"
	case ItemComment:
		return "COMMENT"
	case ItemNumber:
		return "NUMBER"
	case ItemNumberFraction:
		return "FRACTION"
	case ItemSep:
		return "SEPARATOR"
	case ItemUnit:
		return "UNIT"
	case ItemSizeSuffix:
		return "SIZE_SUFFIX"
	case ItemTimeUnit:
		return "TIME_UNIT"
	case ItemTemperatureUnit:
		return "TEMP_UNIT"
	case ItemIdentifierRange:
		return "RANGE"
	case ItemIngredient:
		return "INGREDIENT"
	default:
		return fmt.Sprintf("Unknown [%d]", it)
	}
}

func (it TokenType) Weight() int {
	switch it {
	case ItemNumber, ItemNumberFraction:
		return 10
	case ItemUnit:
		return 20
	case ItemTimeUnit, ItemTemperatureUnit:
		return 15
	case ItemSizeSuffix:
		return 5
	case ItemComment:
		return 2
	case ItemIngredient:
		return 20
	case ItemIdentifierRange:
		return 5
	default:
		return 0
	}
}
