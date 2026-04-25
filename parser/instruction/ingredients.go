package instruction

import (
	"fmt"
	"strings"

	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/lexer"
)

// collectIngredients identifies ingredients using a map of known names.
func collectIngredients(tokens []lexer.Token, text string, refMap map[string]model.IngredientRef, counter *int, prefix string, existing []replacement) ([]model.IngredientRef, []replacement) {
	var ingredients []model.IngredientRef
	var reps []replacement
	seen := make(map[string]bool)

	for _, tok := range tokens {
		if tok.Type != lexer.ItemIngredient {
			continue
		}
		if hasOverlap(tok.StartIndex, tok.EndIndex, existing) {
			continue
		}

		name := strings.ToLower(tok.Lexeme)
		ref, ok := refMap[name]
		if !ok {
			ref = model.IngredientRef{Name: tok.Lexeme}
		}

		// Fill ID if empty
		if ref.ID == "" {
			ref.ID = fmt.Sprintf("%s%d", prefix, *counter)
			*counter++
			// Store the generated ID back in refMap so it's reused for the same name
			refMap[name] = ref
		}

		if !seen[name] {
			ingredients = append(ingredients, ref)
			seen[name] = true
		}

		reps = append(reps, replacement{
			startIdx: tok.StartIndex,
			endIdx:   tok.EndIndex,
			markdown: markdownLink(tok.Lexeme, schemeIngredient, ref.ID),
		})
	}

	return ingredients, reps
}
