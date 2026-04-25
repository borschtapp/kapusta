package instruction

import (
	"testing"

	"github.com/borschtapp/kapusta/parser/lexer"
	"github.com/stretchr/testify/assert"
)

func TestInferIngredients(t *testing.T) {
	text := "Add 1 cup flour and 2 tbsp sugar."
	l, _ := lexer.Lex(text, "en")
	tokens, err := collectTokens(l)
	assert.NoError(t, err)
	counter := 0
	ingredients, reps := inferIngredients(tokens, text, &counter, defaultIDPrefix, nil)
	assert.Len(t, ingredients, 2)
	assert.Len(t, reps, 2)
	assert.Equal(t, "flour", ingredients[0].Name)
	assert.Equal(t, "sugar", ingredients[1].Name)
	assert.Equal(t, "kap0", ingredients[0].ID)
	assert.Equal(t, "kap1", ingredients[1].ID)
}

func TestInferIngredientsNoUnit(t *testing.T) {
	text := "Add salt and pepper."
	l, _ := lexer.Lex(text, "en")
	tokens, err := collectTokens(l)
	assert.NoError(t, err)
	counter := 0
	ingredients, reps := inferIngredients(tokens, text, &counter, defaultIDPrefix, nil)
	assert.Empty(t, ingredients)
	assert.Empty(t, reps)
}

func TestInferIngredientsDeduplicated(t *testing.T) {
	text := "Mix 1 cup flour with 2 cups flour."
	l, _ := lexer.Lex(text, "en")
	tokens, err := collectTokens(l)
	assert.NoError(t, err)
	counter := 0
	ingredients, reps := inferIngredients(tokens, text, &counter, defaultIDPrefix, nil)
	assert.Len(t, ingredients, 1)
	assert.Len(t, reps, 2) // both occurrences of "flour" linked
	assert.Equal(t, "kap0", ingredients[0].ID)
}
