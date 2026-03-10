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
		{"1 tablespoon hot paper, optional", model.Ingredient{Amount: 1, Unit: "tablespoon", UnitCode: "tablespoon", Name: "hot paper", Description: "optional"}},
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
		{"2-3 Eier", model.Ingredient{Amount: 2, MaxAmount: 3, Name: "Eier"}},
		{"Salz + Pfeffer nach Geschmack", model.Ingredient{Name: "Salz + Pfeffer nach Geschmack"}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := Parse(tt.give, "de")
			assert.NoError(t, err)
			assert.Equal(t, &tt.want, got)
		})
	}
}

func TestItalianIngredients(t *testing.T) {
	tests := []struct {
		give string
		want model.Ingredient
	}{
		{"1 mazzetto di erba cipollina", model.Ingredient{Amount: 1, Unit: "mazzetto", UnitCode: "bunch", Name: "di erba cipollina"}},
		{"2 tazze di acqua tiepida, divisa", model.Ingredient{Amount: 2, Unit: "tazze", UnitCode: "cup", Name: "di acqua tiepida", Description: "divisa"}},
		{"¼ tazza di olio extravergine di oliva", model.Ingredient{Amount: 0.25, Unit: "tazza", UnitCode: "cup", Name: "di olio extravergine di oliva"}},
		{"3 rametti di timo fresco", model.Ingredient{Amount: 3, Unit: "rametti", UnitCode: "sprig", Name: "di timo fresco"}},
		{"1 bustina di lievito secco", model.Ingredient{Amount: 1, Unit: "bustina", UnitCode: "pack", Name: "di lievito secco"}},
		{"5 ½ tazze di farina", model.Ingredient{Amount: 5.5, Unit: "tazze", UnitCode: "cup", Name: "di farina"}},
		{"1 cucchiaio di sale marino", model.Ingredient{Amount: 1, Unit: "cucchiaio", UnitCode: "tablespoon", Name: "di sale marino"}},
		{"1 cucchiaino di pepe, opzionale", model.Ingredient{Amount: 1, Unit: "cucchiaino", UnitCode: "teaspoon", Name: "di pepe", Description: "opzionale"}},
		{"2-3 uova", model.Ingredient{Amount: 2, MaxAmount: 3, Name: "uova"}},
		{"500g di carote", model.Ingredient{Amount: 500, Unit: "g", UnitCode: "gram", Name: "di carote"}},
		{"Due cucchiai di maionese", model.Ingredient{Amount: 2, Unit: "cucchiai", UnitCode: "tablespoon", Name: "di maionese"}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := Parse(tt.give, "it")
			assert.NoError(t, err)
			assert.Equal(t, &tt.want, got)
		})
	}
}

func TestUkrainianIngredients(t *testing.T) {
	tests := []struct {
		give string
		want model.Ingredient
	}{
		{"1 пучок цибулі", model.Ingredient{Amount: 1, Unit: "пучок", UnitCode: "bunch", Name: "цибулі"}},
		{"2 мірних чашки теплої води, окремо", model.Ingredient{Amount: 2, Unit: "склянки", UnitCode: "cup", Name: "теплої води", Description: "окремо"}},
		{"¼ мірна чашка оливкової олії першого віджиму", model.Ingredient{Amount: 0.25, Unit: "склянки", UnitCode: "cup", Name: "оливкової олії першого віджиму"}},
		{"3 гілочки свіжого чебрецю", model.Ingredient{Amount: 3, Name: "гілочки свіжого чебрецю"}},
		{"1 пакет сухих дріжджів", model.Ingredient{Amount: 1, Unit: "пакет", UnitCode: "pack", Name: "сухих дріжджів"}},
		{"5 ½ мірних чашок борошна", model.Ingredient{Amount: 5.5, Unit: "склянок", UnitCode: "cup", Name: "борошна"}},
		{"1 столова ложка морської солі", model.Ingredient{Amount: 1, Name: "столова ложка морської солі"}},
		{"1 чайна ложка перцю, за бажанням", model.Ingredient{Amount: 1, Name: "чайна ложка перцю", Description: "за бажанням"}},
		{"2-3 яйця", model.Ingredient{Amount: 2, MaxAmount: 3, Name: "яйця"}},
		{"500г моркви", model.Ingredient{Amount: 500, Unit: "г", UnitCode: "gram", Name: "моркви"}},
		{"Дві столові ложки майонезу", model.Ingredient{Amount: 2, Name: "столові ложки майонезу"}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := Parse(tt.give, "uk")
			assert.NoError(t, err)
			assert.Equal(t, &tt.want, got)
		})
	}
}
