package dictionary_test

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
)

func TestSpanishIngredients(t *testing.T) {
	runIngredientTests(t, "es", []ingredientTestCase{
		{"1 manojo de cebollín", model.Ingredient{Amount: 1, Unit: "manojo", UnitCode: "bunch", Name: "de cebollín"}},
		{"2 tazas de agua tibia, dividida", model.Ingredient{Amount: 2, Unit: "tazas", UnitCode: "cup", Name: "de agua tibia", Description: "dividida"}},
		{"1 cucharada de sal marina", model.Ingredient{Amount: 1, Unit: "cucharada", UnitCode: "tbsp", Name: "de sal marina"}},
		{"3 dientes de ajo", model.Ingredient{Amount: 3, Unit: "dientes", UnitCode: "clove", Name: "de ajo"}},
		{"1/2 cucharadita de pimienta", model.Ingredient{Amount: 0.5, Unit: "cucharadita", UnitCode: "tsp", Name: "de pimienta"}},
		{"500 gramos de harina", model.Ingredient{Amount: 500, Unit: "gramos", UnitCode: "g", Name: "de harina"}},
		{"1 a 2 pizcas de sal", model.Ingredient{Amount: 1, MaxAmount: 2, Unit: "pizcas", UnitCode: "pinch", Name: "de sal"}},
		{"dos rebanadas de pan", model.Ingredient{Amount: 2, Unit: "rebanadas", UnitCode: "slice", Name: "de pan"}},
	})
}

func TestSpanishInstructions(t *testing.T) {
	runInstructionTests(t, "es", []instructionTestCase{
		{
			"Hornee por 15 minutos.",
			nil,
			model.Instruction{
				Text:     "Hornee por 15 minutos.",
				Markdown: `Hornee por [15 minutos](timer:900).`,
				Timers:   []model.Timer{{Value: 900, Raw: "15 minutos"}},
			},
		},
		{
			"Añadir la sal y la pimienta.",
			[]string{"sal", "pimienta"},
			model.Instruction{
				Text:        "Añadir la sal y la pimienta.",
				Markdown:    "Añadir la [sal](ingredient:kap0) y la [pimienta](ingredient:kap1).",
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "sal"}, {ID: "kap1", Name: "pimienta"}},
			},
		},
		{
			"Hornee a 180°C.",
			nil,
			model.Instruction{
				Text:         "Hornee a 180°C.",
				Markdown:     "Hornee a [180°C](temp:180?unit=C).",
				Temperatures: []model.Temperature{{Value: 180, Unit: "C", Raw: "180°C"}},
			},
		},
		{
			"Añada 250g de harina.",
			nil,
			model.Instruction{
				Text:        "Añada 250g de harina.",
				Markdown:    "Añada [250g](amount:250?unit=g) [de](ingredient:kap0) harina.",
				Amounts:     []model.Amount{{Value: 250, Unit: "g", Raw: "250g"}},
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "de"}},
			},
		},
	})
}
