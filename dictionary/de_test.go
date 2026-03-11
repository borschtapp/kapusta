package dictionary_test

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/ingredient"

	"github.com/stretchr/testify/assert"
)

func TestGermanIngredients(t *testing.T) {
	tests := []struct {
		give string
		want model.Ingredient
	}{
		{"Ein Bund Schnittlauch", model.Ingredient{Amount: 1, Unit: "Bund", UnitCode: "bunch", Name: "Schnittlauch"}},
		{"2 Tassen warmes Wasser, geteilt", model.Ingredient{Amount: 2, Unit: "Tassen", UnitCode: "cup", Name: "warmes Wasser", Description: "geteilt"}},
		{"¼ Tasse Olivenöl", model.Ingredient{Amount: 0.25, Unit: "Tasse", UnitCode: "cup", Name: "Olivenöl"}},
		{"3 Zweige Thymian", model.Ingredient{Amount: 3, Unit: "Zweige", UnitCode: "sprig", Name: "Thymian"}},
		{"1 Päckchen Trockenhefe", model.Ingredient{Amount: 1, Unit: "Päckchen", UnitCode: "pack", Name: "Trockenhefe"}},
		{"5 ½ Tassen Mehl", model.Ingredient{Amount: 5.5, Unit: "Tassen", UnitCode: "cup", Name: "Mehl"}},
		{"1 EL Meersalz", model.Ingredient{Amount: 1, Unit: "EL", UnitCode: "tablespoon", Name: "Meersalz"}},
		{"1 EL Pfeffer, optional", model.Ingredient{Amount: 1, Unit: "EL", UnitCode: "tablespoon", Name: "Pfeffer", Description: "optional"}},
		{"2 Stück Gurken", model.Ingredient{Amount: 2, Unit: "Stück", UnitCode: "piece", Name: "Gurken"}},
		{"2-3 Eier", model.Ingredient{Amount: 2, MaxAmount: 3, Name: "Eier"}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := ingredient.Parse(tt.give, "de")
			assert.NoError(t, err)
			assert.Equal(t, &tt.want, got)
		})
	}
}
