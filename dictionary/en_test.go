package dictionary_test

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/ingredient"

	"github.com/stretchr/testify/assert"
)

func TestEnglishIngredients(t *testing.T) {
	tests := []struct {
		give string
		want model.Ingredient
	}{
		{"1 bunch Chives", model.Ingredient{Amount: 1, Unit: "bunch", UnitCode: "bunch", Name: "Chives"}},
		{"2 cups warm water, divided", model.Ingredient{Amount: 2, Unit: "cups", UnitCode: "cup", Name: "warm water", Description: "divided"}},
		{"¼ cup extra-virgin olive oil", model.Ingredient{Amount: 0.25, Unit: "cup", UnitCode: "cup", Name: "extra-virgin olive oil"}},
		{"3 sprigs fresh thyme leaves", model.Ingredient{Amount: 3, Unit: "sprigs", UnitCode: "sprig", Name: "fresh thyme leaves"}},
		{"1 packet active dry yeast (or 2 ¼ teaspoons)", model.Ingredient{Amount: 1, Unit: "packet", UnitCode: "pack", Name: "active dry yeast", Description: "or 2 ¼ teaspoons"}},
		{"5 ½ cups bread flour", model.Ingredient{Amount: 5.5, Unit: "cups", UnitCode: "cup", Name: "bread flour"}},
		{"1 tablespoon sea salt (or kosher salt)", model.Ingredient{Amount: 1, Unit: "tablespoon", UnitCode: "tablespoon", Name: "sea salt", Description: "or kosher salt"}},
		{"1 teaspoon hot paper, optional", model.Ingredient{Amount: 1, Unit: "teaspoon", UnitCode: "teaspoon", Name: "hot paper", Description: "optional"}},
		{"2 pcs. cucumbers", model.Ingredient{Amount: 2, Unit: "pcs.", UnitCode: "piece", Name: "cucumbers"}},
		{"2-3 Pasture-Raised Eggs", model.Ingredient{Amount: 2, MaxAmount: 3, Name: "Pasture-Raised Eggs"}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := ingredient.Parse(tt.give, "en")
			assert.NoError(t, err)
			assert.Equal(t, &tt.want, got)
		})
	}
}
