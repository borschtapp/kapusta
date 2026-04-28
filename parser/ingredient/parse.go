package ingredient

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/lexer"
)

// Options holds optional parameters for ingredient parsing.
type Options struct {
	Lang string
}

func ParseIngredient(str string, opts Options) (model.Ingredient, error) {
	return lexer.Detect(opts.Lang, func(l string) (model.Ingredient, int, error) {
		o := opts
		o.Lang = l
		return parseIngredient(str, o)
	})
}

func parseIngredient(str string, opts Options) (model.Ingredient, int, error) {
	str = strings.ReplaceAll(str, "⁄", "/")

	l, err := lexer.Lex(str, opts.Lang)
	if err != nil {
		return model.Ingredient{}, 0, fmt.Errorf("lex error: %w", err)
	}
	defer l.Close()

	var (
		ing                  model.Ingredient
		unit, name, unitCode string
		prefix               string // fallback name if a unit follows
		secondQty, skipSpace bool
		prev                 lexer.Token
	)

	// Fetch first token
	tok, _ := l.Next()

	for tok.Type != lexer.ItemEOF {
		next, _ := l.Next()

		// 1. Identify context
		isHyphenated := next.Type == lexer.ItemIdentifier && strings.HasPrefix(next.Lexeme, "-")
		isRange := prev.Type == lexer.ItemIdentifierRange
		isMixed := tok.Type == lexer.ItemNumberFraction && prev.Type == lexer.ItemNumber
		isAmount := tok.Type.IsNumber() && !isHyphenated && (ing.Amount == 0 || isRange || isMixed)

		isTextUnit := tok.Type == lexer.ItemUnit && name != "" && (unit != "" || (next.Type == lexer.ItemEOF && !prev.Type.IsNumber()))
		if !isTextUnit && tok.Type == lexer.ItemUnit && name != "" {
			if idx := strings.Index(name, ","); idx > 0 && strings.TrimSpace(name[:idx]) != "" {
				isTextUnit = true
			}
			if !isTextUnit && !prev.Type.IsNumber() && next.Type.IsNumber() && ing.Amount == 0 {
				isTextUnit = true
			}
		}

		// 2. Handle token
		switch {
		case isAmount:
			if isRange {
				ing.MaxAmount = tok.Value
			} else {
				ing.Amount += tok.Value
			}
			skipSpace = false

		case tok.Type == lexer.ItemSep || (tok.Type.IsNumber() && prev.Type == lexer.ItemSep):
			secondQty, skipSpace = true, false

		case tok.Type == lexer.ItemSizeSuffix:
			ing.Description = addDescription(ing.Description, tok.Lexeme)
			skipSpace = false

		case tok.Type == lexer.ItemComment:
			if prev.ShouldMerge(tok) {
				if prev.Type == lexer.ItemUnit || prev.Type == lexer.ItemTimeUnit {
					unit += tok.Lexeme
				} else {
					name += tok.Lexeme
				}
			} else {
				ing.Description = addDescription(ing.Description, lexer.StripParens(tok.Lexeme))
			}
			skipSpace = false

		case tok.Type == lexer.ItemUnit && !secondQty && !isTextUnit:
			if unit == "" && name != "" {
				prefix, name = name, ""
			}
			if unit != "" {
				unit += " "
			}
			unit += tok.Lexeme
			if tok.Code != "other" {
				unitCode = tok.Code
			}
			skipSpace = false

		case tok.Type == lexer.ItemIdentifier || tok.Type == lexer.ItemTimeUnit || tok.Type == lexer.ItemIngredient || isTextUnit || tok.Type.IsNumber():
			name = appendWithSpacing(name, tok.Lexeme, skipSpace, tok.Type)
			skipSpace = tok.Type.IsNumber()
		}

		prev, tok = tok, next
	}

	return finalize(ing, unit, unitCode, name, prefix), l.Score(), nil
}

func finalize(ing model.Ingredient, unit, unitCode, name, prefix string) model.Ingredient {
	// Move trailing descriptions (after commas) from name to description
	if parts := strings.SplitN(name, ",", 2); len(parts) > 1 {
		ing.Description = addDescription(ing.Description, strings.TrimSpace(parts[1]))
		name = strings.TrimSpace(parts[0])
	}

	// Resolve name vs prefix fallback (e.g. "1 garlic clove" -> name="garlic")
	if name == "" {
		name = prefix
	} else if prefix != "" {
		ing.Description = addDescription(prefix, ing.Description)
	}

	ing.Unit, ing.UnitCode, ing.Name = unit, unitCode, name
	return ing
}

func appendWithSpacing(text, lexeme string, skipSpace bool, tokType lexer.TokenType) string {
	if text == "" {
		return lexeme
	}
	shouldAddSpace := !skipSpace
	if skipSpace {
		// Add space after an extra number if the current token is a unit or a non-hyphenated identifier
		if tokType == lexer.ItemUnit || (len(lexeme) > 0 && unicode.IsLetter(rune(lexeme[0])) && !strings.HasPrefix(lexeme, "-")) {
			shouldAddSpace = true
		}
	}

	if shouldAddSpace && len(lexeme) > 0 {
		switch lexeme[0] {
		case ',', '.', ';', ':', '!', '?', ')', ']':
			shouldAddSpace = false
		}
	}

	if shouldAddSpace {
		return text + " " + lexeme
	}
	return text + lexeme
}

func addDescription(existing, new string) string {
	if new == "" {
		return existing
	}
	if existing == "" {
		return new
	}
	return existing + ", " + new
}
