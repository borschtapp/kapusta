package dictionary_test

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/ingredient"

	"github.com/stretchr/testify/assert"
)

func TestUkrainianIngredients(t *testing.T) {
	tests := []struct {
		give string
		want model.Ingredient
	}{
		{"1 пучок цибулі", model.Ingredient{Amount: 1, Unit: "пучок", UnitCode: "bunch", Name: "цибулі"}},
		{"2 мірних чашки теплої води, окремо", model.Ingredient{Amount: 2, Unit: "мірних чашки", UnitCode: "cup", Name: "теплої води", Description: "окремо"}},
		{"¼ мірних чашок оливкової олії першого віджиму", model.Ingredient{Amount: 0.25, Unit: "мірних чашок", UnitCode: "cup", Name: "оливкової олії першого віджиму"}},
		{"3 гілочки свіжого чебрецю", model.Ingredient{Amount: 3, Unit: "гілочки", UnitCode: "sprig", Name: "свіжого чебрецю"}},
		{"1 пакет сухих дріжджів (або 2 ¼ чайних ложки)", model.Ingredient{Amount: 1, Unit: "пакет", UnitCode: "pack", Name: "сухих дріжджів", Description: "або 2 ¼ чайних ложки"}},
		{"5 ½ мірних чашок борошна", model.Ingredient{Amount: 5.5, Unit: "мірних чашок", UnitCode: "cup", Name: "борошна"}},
		{"1 столова ложка морської солі (або кошерна сіль)", model.Ingredient{Amount: 1, Unit: "столова ложка", UnitCode: "tablespoon", Name: "морської солі", Description: "або кошерна сіль"}},
		{"1 чайна ложка перцю, за бажанням", model.Ingredient{Amount: 1, Unit: "чайна ложка", UnitCode: "teaspoon", Name: "перцю", Description: "за бажанням"}},
		{"2 шт. огірки", model.Ingredient{Amount: 2, Unit: "шт.", UnitCode: "piece", Name: "огірки"}},
		{"2-3 яйця", model.Ingredient{Amount: 2, MaxAmount: 3, Name: "яйця"}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := ingredient.Parse(tt.give, "uk")
			assert.NoError(t, err)
			assert.Equal(t, &tt.want, got)
		})
	}
}
