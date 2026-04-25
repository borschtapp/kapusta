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
