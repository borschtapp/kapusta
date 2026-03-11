package dictionary_test

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/ingredient"

	"github.com/stretchr/testify/assert"
)

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
		{"2 pezzi di cetrioli", model.Ingredient{Amount: 2, Unit: "pezzi", UnitCode: "piece", Name: "di cetrioli"}},
		{"2-3 uova", model.Ingredient{Amount: 2, MaxAmount: 3, Name: "uova"}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := ingredient.Parse(tt.give, "it")
			assert.NoError(t, err)
			assert.Equal(t, &tt.want, got)
		})
	}
}
