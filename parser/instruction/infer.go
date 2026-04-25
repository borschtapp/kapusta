package instruction

import (
	"fmt"
	"strings"

	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/lexer"
)

// inferIngredients attempts to find ingredient-like spans that were not matched by collectIngredients.
// It follows the heuristic: identifiers following a quantity (number + unit) are likely ingredients.
func inferIngredients(tokens []lexer.Token, text string, idCounter *int, idPrefix string, existing []replacement) ([]model.IngredientRef, []replacement) {
	var ingreds []model.IngredientRef
	var reps []replacement
	seen := make(map[string]string) // name -> id

	c := &lexer.TokenCursor{Tokens: tokens}

	for !c.IsEOF() {
		tok := c.Peek()
		if tok.Type.IsNumber() {
			parseQuantity(c)
			if next := c.Peek(); next.Type == lexer.ItemUnit {
				c.Next()
				if after := c.Peek(); after.Type == lexer.ItemIdentifier {
					if hasOverlap(after.StartIndex, after.EndIndex, existing) {
						c.Next()
						continue
					}
					name := strings.TrimSuffix(after.Lexeme, ".")
					nameLower := strings.ToLower(name)
					id, exists := seen[nameLower]
					if !exists {
						id = fmt.Sprintf("%s%d", idPrefix, *idCounter)
						*idCounter++
						seen[nameLower] = id
						ingreds = append(ingreds, model.IngredientRef{
							ID:   id,
							Name: name,
						})
					}

					endIdx := after.EndIndex
					if strings.HasSuffix(after.Lexeme, ".") {
						endIdx--
					}

					reps = append(reps, replacement{
						startIdx: after.StartIndex,
						endIdx:   endIdx,
						markdown: markdownLink(name, schemeIngredient, id),
					})
					c.Next()
					continue
				}
			}
		}
		c.Next()
	}

	return ingreds, reps
}
