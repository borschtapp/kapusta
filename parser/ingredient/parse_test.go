package ingredient

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/borschtapp/krip/model"
)

func TestParse(t *testing.T) {
	ing, err := Parse("2 cups warm water divided", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(2), ing.Quantity)
	assert.Equal(t, "cups", ing.Unit)
	assert.Equal(t, "warm water divided", ing.Ingredient)
}

func TestParseRange(t *testing.T) {
	ing, err := Parse("1-2 cups blueberries ((add after cooking base recipe))", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(1), ing.Quantity)
	assert.Equal(t, float64(2), ing.QuantityMax)
	assert.Equal(t, "cups", ing.Unit)
	assert.Equal(t, "blueberries", ing.Ingredient)
	assert.Equal(t, "add after cooking base recipe", ing.Annotation)
}

func TestParseMultipleUnits(t *testing.T) {
	ing, err := Parse("2 cups / 480 ml water", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(2), ing.Quantity)
	assert.Equal(t, "cups", ing.Unit)
	assert.Equal(t, "water", ing.Ingredient)
}

func TestParseMergedUnit(t *testing.T) {
	ing, err := Parse("500g carrots", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(500), ing.Quantity)
	assert.Equal(t, "g", ing.Unit)
	assert.Equal(t, "carrots", ing.Ingredient)
}

func TestParseMergedUnit2(t *testing.T) {
	ing, err := Parse("2EL Mayonnaise", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(2), ing.Quantity)
	assert.Empty(t, ing.Unit)
	assert.Equal(t, "EL Mayonnaise", ing.Ingredient)
}

func TestParseInches(t *testing.T) {
	ing, err := Parse("4 med. potatoes, pared & cut lengthwise, in 1/4\" slices", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(4), ing.Quantity)
	assert.Empty(t, ing.Unit)
	assert.Equal(t, "med. potatoes", ing.Ingredient)
	assert.Equal(t, "pared & cut lengthwise, in 1/4\" slices", ing.Annotation)
}

func TestMultiple(t *testing.T) {
	tests := []struct {
		give string
		want model.Ingredient
	}{
		{"1 bunch Chives", model.Ingredient{Quantity: 1, Unit: "bunch", Ingredient: "Chives"}},
		{"2 cups warm water, divided", model.Ingredient{Quantity: 2, Unit: "cups", Ingredient: "warm water", Annotation: "divided"}},
		{"¼ cup extra-virgin olive oil", model.Ingredient{Quantity: 0.25, Unit: "cup", Ingredient: "extra-virgin olive oil"}},
		{"3 sprigs fresh thyme leaves", model.Ingredient{Quantity: 3, Unit: "sprigs", Ingredient: "fresh thyme leaves"}},
		{"1 packet active dry yeast (or 2 ¼ teaspoons)", model.Ingredient{Quantity: 1, Unit: "packet", Ingredient: "active dry yeast", Annotation: "or 2 ¼ teaspoons"}},
		{"5 ½ cups bread flour", model.Ingredient{Quantity: 5.5, Unit: "cups", Ingredient: "bread flour"}},
		{"1 tablespoon sea salt (or kosher salt)", model.Ingredient{Quantity: 1, Unit: "tablespoon", Ingredient: "sea salt", Annotation: "or kosher salt"}},
		{"1 tablespoon hot paper, optional", model.Ingredient{Quantity: 1, Unit: "tablespoon", Ingredient: "hot paper", Annotation: "optional"}},
		{"2 Persian Cucumbers", model.Ingredient{Quantity: 2, Ingredient: "Persian Cucumbers"}},
		{"2-3 Pasture-Raised Eggs", model.Ingredient{Quantity: 2, QuantityMax: 3, Ingredient: "Pasture-Raised Eggs"}},
		{"Cinnamon + Sea salt to taste", model.Ingredient{Ingredient: "Cinnamon + Sea salt to taste"}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := Parse(tt.give, "en")
			assert.NoError(t, err)
			assert.Equal(t, &tt.want, got)
		})
	}
}
