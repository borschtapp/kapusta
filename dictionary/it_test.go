package dictionary_test

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
)

func TestItalianIngredients(t *testing.T) {
	runIngredientTests(t, "it", []ingredientTestCase{
		{"1 mazzetto di erba cipollina", model.Ingredient{Amount: 1, Unit: "mazzetto", UnitCode: "bunch", Name: "di erba cipollina"}},
		{"2 tazze di acqua tiepida, divisa", model.Ingredient{Amount: 2, Unit: "tazze", UnitCode: "cup", Name: "di acqua tiepida", Description: "divisa"}},
		{"¼ tazza di olio extravergine di oliva", model.Ingredient{Amount: 0.25, Unit: "tazza", UnitCode: "cup", Name: "di olio extravergine di oliva"}},
		{"3 rametti di timo fresco", model.Ingredient{Amount: 3, Unit: "rametti", UnitCode: "sprig", Name: "di timo fresco"}},
		{"1 bustina di lievito secco", model.Ingredient{Amount: 1, Unit: "bustina", UnitCode: "pkg", Name: "di lievito secco"}},
		{"5 ½ tazze di farina", model.Ingredient{Amount: 5.5, Unit: "tazze", UnitCode: "cup", Name: "di farina"}},
		{"1 cucchiaio di sale marino", model.Ingredient{Amount: 1, Unit: "cucchiaio", UnitCode: "tbsp", Name: "di sale marino"}},
		{"1 cucchiaino di pepe, opzionale", model.Ingredient{Amount: 1, Unit: "cucchiaino", UnitCode: "tsp", Name: "di pepe", Description: "opzionale"}},
		{"2 pezzi di cetrioli", model.Ingredient{Amount: 2, Unit: "pezzi", UnitCode: "pc", Name: "di cetrioli"}},
		{"2-3 uova", model.Ingredient{Amount: 2, MaxAmount: 3, Name: "uova"}},
	})
}

func TestItalianInstructions(t *testing.T) {
	runInstructionTests(t, "it", []instructionTestCase{
		{
			"Cuocere in forno per 15 minuti.",
			nil,
			model.Instruction{
				Text:     "Cuocere in forno per 15 minuti.",
				Markdown: `Cuocere in forno per [15 minuti](timer:900).`,
				Timers:   []model.Timer{{Value: 900, Raw: "15 minuti"}},
			},
		},
		{
			"Aggiungere sale e pepe.",
			[]string{"sale", "pepe"},
			model.Instruction{
				Text:        "Aggiungere sale e pepe.",
				Markdown:    "Aggiungere [sale](ingredient:kap0) e [pepe](ingredient:kap1).",
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "sale"}, {ID: "kap1", Name: "pepe"}},
			},
		},
		{
			"Cuocere a 180°C.",
			nil,
			model.Instruction{
				Text:         "Cuocere a 180°C.",
				Markdown:     "Cuocere a [180°C](temp:180?unit=C).",
				Temperatures: []model.Temperature{{Value: 180, Unit: "C", Raw: "180°C"}},
			},
		},
		{
			"Aggiungere 250g di farina.",
			nil,
			model.Instruction{
				Text:        "Aggiungere 250g di farina.",
				Markdown:    "Aggiungere [250g](amount:250?unit=g) [di](ingredient:kap0) farina.",
				Amounts:     []model.Amount{{Value: 250, Unit: "g", Raw: "250g"}},
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "di"}},
			},
		},
	})
}
