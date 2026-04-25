package dictionary_test

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
)

func TestFrenchIngredients(t *testing.T) {
	runIngredientTests(t, "fr", []ingredientTestCase{
		{"1 botte de ciboulette", model.Ingredient{Amount: 1, Unit: "botte", UnitCode: "bunch", Name: "de ciboulette"}},
		{"2 tasses d'eau tiède, divisée", model.Ingredient{Amount: 2, Unit: "tasses", UnitCode: "cup", Name: "d'eau tiède", Description: "divisée"}},
		{"1 cuillère à soupe de sel", model.Ingredient{Amount: 1, Unit: "cuillère à soupe", UnitCode: "tbsp", Name: "de sel"}},
		{"3 gousses d'ail", model.Ingredient{Amount: 3, Unit: "gousses", UnitCode: "clove", Name: "d'ail"}},
		{"1/2 cuillère à café de poivre", model.Ingredient{Amount: 0.5, Unit: "cuillère à café", UnitCode: "tsp", Name: "de poivre"}},
		{"500 grammes de farine", model.Ingredient{Amount: 500, Unit: "grammes", UnitCode: "g", Name: "de farine"}},
		{"1 à 2 pincées de sel", model.Ingredient{Amount: 1, MaxAmount: 2, Unit: "pincées", UnitCode: "pinch", Name: "de sel"}},
		{"deux tranches de pain", model.Ingredient{Amount: 2, Unit: "tranches", UnitCode: "slice", Name: "de pain"}},
	})
}

func TestFrenchInstructions(t *testing.T) {
	runInstructionTests(t, "fr", []instructionTestCase{
		{
			"Cuire au four pendant 15 minutes.",
			nil,
			model.Instruction{
				Text:     "Cuire au four pendant 15 minutes.",
				Markdown: `Cuire au four pendant [15 minutes](timer:900).`,
				Timers:   []model.Timer{{Value: 900, Raw: "15 minutes"}},
			},
		},
		{
			"Ajoutez le sel et le poivre.",
			[]string{"sel", "poivre"},
			model.Instruction{
				Text:        "Ajoutez le sel et le poivre.",
				Markdown:    "Ajoutez le [sel](ingredient:kap0) et le [poivre](ingredient:kap1).",
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "sel"}, {ID: "kap1", Name: "poivre"}},
			},
		},
		{
			"Cuire à 180°C.",
			nil,
			model.Instruction{
				Text:         "Cuire à 180°C.",
				Markdown:     "Cuire à [180°C](temp:180?unit=C).",
				Temperatures: []model.Temperature{{Value: 180, Unit: "C", Raw: "180°C"}},
			},
		},
		{
			"Ajouter 250g de farine.",
			nil,
			model.Instruction{
				Text:        "Ajouter 250g de farine.",
				Markdown:    "Ajouter [250g](amount:250?unit=g) [de](ingredient:kap0) farine.",
				Amounts:     []model.Amount{{Value: 250, Unit: "g", Raw: "250g"}},
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "de"}},
			},
		},
	})
}
