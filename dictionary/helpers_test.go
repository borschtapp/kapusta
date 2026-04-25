package dictionary_test

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/ingredient"
	"github.com/stretchr/testify/assert"
)

type ingredientTestCase struct {
	give string
	want model.Ingredient
}

func runIngredientTests(t *testing.T, lang string, tests []ingredientTestCase) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := ingredient.Parse(tt.give, ingredient.Options{Lang: lang})
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
