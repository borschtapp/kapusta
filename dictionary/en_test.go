package dictionary_test

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
)

func TestEnglishIngredients(t *testing.T) {
	runIngredientTests(t, "en", []ingredientTestCase{
		{"1 bunch Chives", model.Ingredient{Amount: 1, Unit: "bunch", UnitCode: "bunch", Name: "Chives"}},
		{"2 cups warm water, divided", model.Ingredient{Amount: 2, Unit: "cups", UnitCode: "cup", Name: "warm water", Description: "divided"}},
		{"¼ cup extra-virgin olive oil", model.Ingredient{Amount: 0.25, Unit: "cup", UnitCode: "cup", Name: "extra-virgin olive oil"}},
		{"3 sprigs fresh thyme leaves", model.Ingredient{Amount: 3, Unit: "sprigs", UnitCode: "sprig", Name: "fresh thyme leaves"}},
		{"1 packet active dry yeast (or 2 ¼ teaspoons)", model.Ingredient{Amount: 1, Unit: "packet", UnitCode: "pkg", Name: "active dry yeast", Description: "or 2 ¼ teaspoons"}},
		{"5 ½ cups bread flour", model.Ingredient{Amount: 5.5, Unit: "cups", UnitCode: "cup", Name: "bread flour"}},
		{"1 tablespoon sea salt (or kosher salt)", model.Ingredient{Amount: 1, Unit: "tablespoon", UnitCode: "tbsp", Name: "sea salt", Description: "or kosher salt"}},
		{"1 teaspoon hot paper, optional", model.Ingredient{Amount: 1, Unit: "teaspoon", UnitCode: "tsp", Name: "hot paper", Description: "optional"}},
		{"2 pcs. cucumbers", model.Ingredient{Amount: 2, Unit: "pcs.", UnitCode: "pc", Name: "cucumbers"}},
		{"2-3 Pasture-Raised Eggs", model.Ingredient{Amount: 2, MaxAmount: 3, Name: "Pasture-Raised Eggs"}},
		{"2 medium Eggs", model.Ingredient{Amount: 2, Name: "Eggs", Description: "medium"}},
	})
}

func TestEnglishInstructions(t *testing.T) {
	runInstructionTests(t, "en", []instructionTestCase{
		{
			"Bake for 15 minutes.",
			nil,
			model.Instruction{
				Text:     "Bake for 15 minutes.",
				Markdown: `Bake for [15 minutes](timer:900).`,
				Timers:   []model.Timer{{Value: 900, Raw: "15 minutes"}},
			},
		},
		{
			"Simmer for 2-3 hrs.",
			nil,
			model.Instruction{
				Text:     "Simmer for 2-3 hrs.",
				Markdown: `Simmer for [2-3 hrs](timer:7200?max=10800).`,
				Timers:   []model.Timer{{Value: 7200, MaxValue: 10800, Raw: "2-3 hrs"}},
			},
		},
		{
			"Add salt and pepper.",
			[]string{"salt", "pepper"},
			model.Instruction{
				Text:        "Add salt and pepper.",
				Markdown:    "Add [salt](ingredient:kap0) and [pepper](ingredient:kap1).",
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "salt"}, {ID: "kap1", Name: "pepper"}},
			},
		},
		{
			"Bake at 350°F.",
			nil,
			model.Instruction{
				Text:         "Bake at 350°F.",
				Markdown:     "Bake at [350°F](temp:350?unit=F).",
				Temperatures: []model.Temperature{{Value: 350, Unit: "F", Raw: "350°F"}},
			},
		},
		{
			"Add 2 cups flour.",
			nil,
			model.Instruction{
				Text:        "Add 2 cups flour.",
				Markdown:    "Add [2 cups](amount:2?unit=cup) [flour](ingredient:kap0).",
				Amounts:     []model.Amount{{Value: 2, Unit: "cup", Raw: "2 cups"}},
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "flour"}},
			},
		},
	})
}
