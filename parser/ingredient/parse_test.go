package ingredient

import (
	"testing"

	"github.com/borschtapp/kapusta/model"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	ing, err := Parse("2 cups warm water divided", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(2), ing.Amount)
	assert.Equal(t, "cups", ing.Unit)
	assert.Equal(t, "warm water divided", ing.Name)
}

func TestParseRange(t *testing.T) {
	ing, err := Parse("1-2 cups blueberries ((add after cooking base recipe))", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(1), ing.Amount)
	assert.Equal(t, float64(2), ing.MaxAmount)
	assert.Equal(t, "cups", ing.Unit)
	assert.Equal(t, "blueberries", ing.Name)
	assert.Equal(t, "add after cooking base recipe", ing.Description)
}

func TestParseMultipleUnits(t *testing.T) {
	ing, err := Parse("2 cups / 480 ml water", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(2), ing.Amount)
	assert.Equal(t, "cups", ing.Unit)
	assert.Equal(t, "water", ing.Name)
}

func TestParseMergedUnit(t *testing.T) {
	ing, err := Parse("500g carrots", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(500), ing.Amount)
	assert.Equal(t, "g", ing.Unit)
	assert.Equal(t, "carrots", ing.Name)
}

func TestParseMergedUnit2(t *testing.T) {
	ing, err := Parse("2EL Mayonnaise", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(2), ing.Amount)
	assert.Empty(t, ing.Unit)
	assert.Equal(t, "EL Mayonnaise", ing.Name)
}

func TestParseInches(t *testing.T) {
	ing, err := Parse("4 med. potatoes, pared & cut lengthwise, in 1/4\" slices", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(4), ing.Amount)
	assert.Empty(t, ing.Unit)
	assert.Equal(t, "med. potatoes", ing.Name)
	assert.Equal(t, "pared & cut lengthwise, in 1/4\" slices", ing.Description)
}

func TestMultiple(t *testing.T) {
	tests := []struct {
		give string
		want model.Ingredient
	}{
		{"1 bunch Chives", model.Ingredient{Amount: 1, Unit: "bunch", Name: "Chives"}},
		{"2 cups warm water, divided", model.Ingredient{Amount: 2, Unit: "cups", Name: "warm water", Description: "divided"}},
		{"¼ cup extra-virgin olive oil", model.Ingredient{Amount: 0.25, Unit: "cup", Name: "extra-virgin olive oil"}},
		{"3 sprigs fresh thyme leaves", model.Ingredient{Amount: 3, Unit: "sprigs", Name: "fresh thyme leaves"}},
		{"1 packet active dry yeast (or 2 ¼ teaspoons)", model.Ingredient{Amount: 1, Unit: "packet", Name: "active dry yeast", Description: "or 2 ¼ teaspoons"}},
		{"5 ½ cups bread flour", model.Ingredient{Amount: 5.5, Unit: "cups", Name: "bread flour"}},
		{"1 tablespoon sea salt (or kosher salt)", model.Ingredient{Amount: 1, Unit: "tablespoon", Name: "sea salt", Description: "or kosher salt"}},
		{"1 tablespoon hot paper, optional", model.Ingredient{Amount: 1, Unit: "tablespoon", Name: "hot paper", Description: "optional"}},
		{"2 Persian Cucumbers", model.Ingredient{Amount: 2, Name: "Persian Cucumbers"}},
		{"2-3 Pasture-Raised Eggs", model.Ingredient{Amount: 2, MaxAmount: 3, Name: "Pasture-Raised Eggs"}},
		{"Cinnamon + Sea salt to taste", model.Ingredient{Name: "Cinnamon + Sea salt to taste"}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := Parse(tt.give, "en")
			assert.NoError(t, err)
			assert.Equal(t, &tt.want, got)
		})
	}
}
