package ingredient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	ing, err := Parse("2 cups warm water divided", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(2), ing.Amount)
	assert.Equal(t, "cups", ing.Unit)
	assert.Equal(t, "cup", ing.UnitCode)
	assert.Equal(t, "warm water divided", ing.Name)
}

func TestParseRange(t *testing.T) {
	ing, err := Parse("1-2 cups blueberries ((add after cooking base recipe))", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(1), ing.Amount)
	assert.Equal(t, float64(2), ing.MaxAmount)
	assert.Equal(t, "cups", ing.Unit)
	assert.Equal(t, "cup", ing.UnitCode)
	assert.Equal(t, "blueberries", ing.Name)
	assert.Equal(t, "add after cooking base recipe", ing.Description)
}

func TestParseMultipleUnits(t *testing.T) {
	ing, err := Parse("2 cups / 480 ml water", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(2), ing.Amount)
	assert.Equal(t, "cups", ing.Unit)
	assert.Equal(t, "cup", ing.UnitCode)
	assert.Equal(t, "water", ing.Name)
}

func TestParseMergedUnit(t *testing.T) {
	ing, err := Parse("500g carrots", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(500), ing.Amount)
	assert.Equal(t, "g", ing.Unit)
	assert.Equal(t, "gram", ing.UnitCode)
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

func TestParseSlashFraction(t *testing.T) {
	ing, err := Parse("1/2 cup milk", "en")
	assert.NoError(t, err)
	assert.Equal(t, 0.5, ing.Amount)
	assert.Equal(t, "cup", ing.Unit)
	assert.Equal(t, "cup", ing.UnitCode)
	assert.Equal(t, "milk", ing.Name)
}

func TestParseUnicodeFractionSlash(t *testing.T) {
	ing, err := Parse("1\u20442 cup milk", "en")
	assert.NoError(t, err)
	assert.Equal(t, 0.5, ing.Amount)
	assert.Equal(t, "cup", ing.Unit)
	assert.Equal(t, "milk", ing.Name)
}

func TestParseMixedNumberSlash(t *testing.T) {
	ing, err := Parse("1 1/2 cups flour", "en")
	assert.NoError(t, err)
	assert.Equal(t, 1.5, ing.Amount)
	assert.Equal(t, "cups", ing.Unit)
	assert.Equal(t, "flour", ing.Name)
}

func TestParseDecimalAmount(t *testing.T) {
	ing, err := Parse("1.5 tablespoons olive oil", "en")
	assert.NoError(t, err)
	assert.Equal(t, 1.5, ing.Amount)
	assert.Equal(t, "tablespoons", ing.Unit)
	assert.Equal(t, "tablespoon", ing.UnitCode)
	assert.Equal(t, "olive oil", ing.Name)
}

func TestParseWordNumber(t *testing.T) {
	ing, err := Parse("two cups flour", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(2), ing.Amount)
	assert.Equal(t, "cups", ing.Unit)
	assert.Equal(t, "cup", ing.UnitCode)
	assert.Equal(t, "flour", ing.Name)
}

func TestParseArticleAsOne(t *testing.T) {
	ing, err := Parse("a pinch of salt", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(1), ing.Amount)
	assert.Equal(t, "pinch", ing.Unit)
	assert.Equal(t, "of salt", ing.Name)
}

func TestParseWordRange(t *testing.T) {
	ing, err := Parse("1 to 2 cups sugar", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(1), ing.Amount)
	assert.Equal(t, float64(2), ing.MaxAmount)
	assert.Equal(t, "cups", ing.Unit)
	assert.Equal(t, "sugar", ing.Name)
}

func TestParseCaseInsensitiveUnit(t *testing.T) {
	ing, err := Parse("2 CUPS water", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(2), ing.Amount)
	assert.Equal(t, "cup", ing.UnitCode)
	assert.Equal(t, "water", ing.Name)
}

func TestParseNoAmount(t *testing.T) {
	ing, err := Parse("salt", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(0), ing.Amount)
	assert.Empty(t, ing.Unit)
	assert.Equal(t, "salt", ing.Name)
}

func TestParseNoUnit(t *testing.T) {
	ing, err := Parse("3 eggs", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(3), ing.Amount)
	assert.Empty(t, ing.Unit)
	assert.Equal(t, "eggs", ing.Name)
}

func TestParseMultipleCommas(t *testing.T) {
	ing, err := Parse("flour, organic, finely sifted", "en")
	assert.NoError(t, err)
	assert.Equal(t, "flour", ing.Name)
	assert.Equal(t, "organic, finely sifted", ing.Description)
}

func TestParsePlusAsSeparator(t *testing.T) {
	ing, err := Parse("salt + pepper to taste", "en")
	assert.NoError(t, err)
	assert.Equal(t, float64(0), ing.Amount)
	assert.Equal(t, "salt + pepper to taste", ing.Name)
}
