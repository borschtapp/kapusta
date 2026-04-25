package dictionary_test

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
)

func TestGermanIngredients(t *testing.T) {
	runIngredientTests(t, "de", []ingredientTestCase{
		{"Ein Bund Schnittlauch", model.Ingredient{Amount: 1, Unit: "Bund", UnitCode: "bunch", Name: "Schnittlauch"}},
		{"2 Tassen warmes Wasser, geteilt", model.Ingredient{Amount: 2, Unit: "Tassen", UnitCode: "cup", Name: "warmes Wasser", Description: "geteilt"}},
		{"¼ Tasse Olivenöl", model.Ingredient{Amount: 0.25, Unit: "Tasse", UnitCode: "cup", Name: "Olivenöl"}},
		{"3 Zweige Thymian", model.Ingredient{Amount: 3, Unit: "Zweige", UnitCode: "sprig", Name: "Thymian"}},
		{"1 Päckchen Trockenhefe", model.Ingredient{Amount: 1, Unit: "Päckchen", UnitCode: "pkg", Name: "Trockenhefe"}},
		{"5 ½ Tassen Mehl", model.Ingredient{Amount: 5.5, Unit: "Tassen", UnitCode: "cup", Name: "Mehl"}},
		{"1 EL Meersalz", model.Ingredient{Amount: 1, Unit: "EL", UnitCode: "tbsp", Name: "Meersalz"}},
		{"1 EL Pfeffer, optional", model.Ingredient{Amount: 1, Unit: "EL", UnitCode: "tbsp", Name: "Pfeffer", Description: "optional"}},
		{"2 Stück Gurken", model.Ingredient{Amount: 2, Unit: "Stück", UnitCode: "pc", Name: "Gurken"}},
		{"2-3 Eier", model.Ingredient{Amount: 2, MaxAmount: 3, Name: "Eier"}},
	})
}

func TestGermanInstructions(t *testing.T) {
	runInstructionTests(t, "de", []instructionTestCase{
		{
			"15 Minuten backen.",
			nil,
			model.Instruction{
				Text:     "15 Minuten backen.",
				Markdown: "[15 Minuten](timer:900) backen.",
				Timers:   []model.Timer{{Value: 900, Raw: "15 Minuten"}},
			},
		},
		{
			"Kochen Sie es für 2-3 Min.",
			nil,
			model.Instruction{
				Text:     "Kochen Sie es für 2-3 Min.",
				Markdown: `Kochen Sie es für [2-3 Min](timer:120?max=180).`,
				Timers:   []model.Timer{{Value: 120, MaxValue: 180, Raw: "2-3 Min"}},
			},
		},
		{
			"Mit Salz und Pfeffer würzen.",
			[]string{"Salz", "Pfeffer"},
			model.Instruction{
				Text:        "Mit Salz und Pfeffer würzen.",
				Markdown:    "Mit [Salz](ingredient:kap0) und [Pfeffer](ingredient:kap1) würzen.",
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "Salz"}, {ID: "kap1", Name: "Pfeffer"}},
			},
		},
		{
			"Bei 180°C backen.",
			nil,
			model.Instruction{
				Text:         "Bei 180°C backen.",
				Markdown:     "Bei [180°C](temp:180?unit=C) backen.",
				Temperatures: []model.Temperature{{Value: 180, Unit: "C", Raw: "180°C"}},
			},
		},
		{
			"250g Mehl hinzufügen.",
			nil,
			model.Instruction{
				Text:        "250g Mehl hinzufügen.",
				Markdown:    "[250g](amount:250?unit=g) [Mehl](ingredient:kap0) hinzufügen.",
				Amounts:     []model.Amount{{Value: 250, Unit: "g", Raw: "250g"}},
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "Mehl"}},
			},
		},
	})
}
