package dictionary_test

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/ingredient"

	"github.com/stretchr/testify/assert"
)

func TestFrenchIngredients(t *testing.T) {
	tests := []struct {
		give string
		want model.Ingredient
	}{
		{"1 botte de ciboulette", model.Ingredient{Amount: 1, Unit: "botte", UnitCode: "bunch", Name: "de ciboulette"}},
		{"2 tasses d'eau tiède, divisée", model.Ingredient{Amount: 2, Unit: "tasses", UnitCode: "cup", Name: "d'eau tiède", Description: "divisée"}},
		{"1 cuillère à soupe de sel", model.Ingredient{Amount: 1, Unit: "cuillère à soupe", UnitCode: "tbsp", Name: "de sel"}},
		{"3 gousses d'ail", model.Ingredient{Amount: 3, Unit: "gousses", UnitCode: "clove", Name: "d'ail"}},
		{"1/2 cuillère à café de poivre", model.Ingredient{Amount: 0.5, Unit: "cuillère à café", UnitCode: "tsp", Name: "de poivre"}},
		{"500 grammes de farine", model.Ingredient{Amount: 500, Unit: "grammes", UnitCode: "g", Name: "de farine"}},
		{"1 à 2 pincées de sel", model.Ingredient{Amount: 1, MaxAmount: 2, Unit: "pincées", UnitCode: "pinch", Name: "de sel"}},
		{"deux tranches de pain", model.Ingredient{Amount: 2, Unit: "tranches", UnitCode: "slice", Name: "de pain"}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := ingredient.Parse(tt.give, "fr")
			assert.NoError(t, err)
			assert.Equal(t, &tt.want, got)
		})
	}
}
