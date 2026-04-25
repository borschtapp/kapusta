package dictionary_test

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
)

func TestUkrainianIngredients(t *testing.T) {
	runIngredientTests(t, "uk", []ingredientTestCase{
		{"1 пучок цибулі", model.Ingredient{Amount: 1, Unit: "пучок", UnitCode: "bunch", Name: "цибулі"}},
		{"2 мірних чашки теплої води, окремо", model.Ingredient{Amount: 2, Unit: "мірних чашки", UnitCode: "cup", Name: "теплої води", Description: "окремо"}},
		{"¼ мірних чашок оливкової олії першого віджиму", model.Ingredient{Amount: 0.25, Unit: "мірних чашок", UnitCode: "cup", Name: "оливкової олії першого віджиму"}},
		{"3 гілочки свіжого чебрецю", model.Ingredient{Amount: 3, Unit: "гілочки", UnitCode: "sprig", Name: "свіжого чебрецю"}},
		{"1 пакет сухих дріжджів (або 2 ¼ чайних ложки)", model.Ingredient{Amount: 1, Unit: "пакет", UnitCode: "pkg", Name: "сухих дріжджів", Description: "або 2 ¼ чайних ложки"}},
		{"5 ½ мірних чашок борошна", model.Ingredient{Amount: 5.5, Unit: "мірних чашок", UnitCode: "cup", Name: "борошна"}},
		{"1 столова ложка морської солі (або кошерна сіль)", model.Ingredient{Amount: 1, Unit: "столова ложка", UnitCode: "tbsp", Name: "морської солі", Description: "або кошерна сіль"}},
		{"1 чайна ложка перцю, за бажанням", model.Ingredient{Amount: 1, Unit: "чайна ложка", UnitCode: "tsp", Name: "перцю", Description: "за бажанням"}},
		{"2 шт. огірки", model.Ingredient{Amount: 2, Unit: "шт.", UnitCode: "pc", Name: "огірки"}},
		{"2-3 яйця", model.Ingredient{Amount: 2, MaxAmount: 3, Name: "яйця"}},
	})
}

func TestUkrainianInstructions(t *testing.T) {
	runInstructionTests(t, "uk", []instructionTestCase{
		{
			"Варіть 15 хвилин.",
			nil,
			model.Instruction{
				Text:     "Варіть 15 хвилин.",
				Markdown: `Варіть [15 хвилин](timer:900).`,
				Timers:   []model.Timer{{Value: 900, Raw: "15 хвилин"}},
			},
		},
		{
			"Смажте 2-3 хв.",
			nil,
			model.Instruction{
				Text:     "Смажте 2-3 хв.",
				Markdown: `Смажте [2-3 хв](timer:120?max=180).`,
				Timers:   []model.Timer{{Value: 120, MaxValue: 180, Raw: "2-3 хв"}},
			},
		},
		{
			"Додайте сіль та перець.",
			[]string{"сіль", "перець"},
			model.Instruction{
				Text:        "Додайте сіль та перець.",
				Markdown:    "Додайте [сіль](ingredient:kap0) та [перець](ingredient:kap1).",
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "сіль"}, {ID: "kap1", Name: "перець"}},
			},
		},
		{
			"Випікати при 180°C.",
			nil,
			model.Instruction{
				Text:         "Випікати при 180°C.",
				Markdown:     "Випікати при [180°C](temp:180?unit=C).",
				Temperatures: []model.Temperature{{Value: 180, Unit: "C", Raw: "180°C"}},
			},
		},
		{
			"Додайте 250г борошна.",
			nil,
			model.Instruction{
				Text:        "Додайте 250г борошна.",
				Markdown:    "Додайте [250г](amount:250?unit=g) [борошна](ingredient:kap0).",
				Amounts:     []model.Amount{{Value: 250, Unit: "g", Raw: "250г"}},
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "борошна"}},
			},
		},
	})
}
