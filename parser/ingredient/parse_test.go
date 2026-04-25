package ingredient

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
	"github.com/stretchr/testify/assert"
)

func TestParseIngredient(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  model.Ingredient
	}{
		{
			name:  "nested brackets",
			input: "2 cups flour (sifted (very fine) thoroughly)",
			want:  model.Ingredient{Amount: 2, Unit: "cups", UnitCode: "cup", Name: "flour", Description: "sifted (very fine) thoroughly"},
		},
		{
			name:  "basic",
			input: "2 cups warm water divided",
			want:  model.Ingredient{Amount: 2, Unit: "cups", UnitCode: "cup", Name: "warm water divided"},
		},
		{
			name:  "range",
			input: "1-2 cups blueberries ((add after cooking base recipe))",
			want:  model.Ingredient{Amount: 1, MaxAmount: 2, Unit: "cups", UnitCode: "cup", Name: "blueberries", Description: "add after cooking base recipe"},
		},
		{
			name:  "unit with size",
			input: "1 10-ounce bag frozen cherries",
			want:  model.Ingredient{Amount: 1, Unit: "bag", UnitCode: "", Name: "frozen cherries", Description: "10-ounce"},
		},
		{
			name:  "unit with size and comma",
			input: "2 1½-inch-thick bone-in pork rib chops (about 1 pound each), patted dry",
			want:  model.Ingredient{Amount: 2, Unit: "", UnitCode: "", Name: "1½-inch-thick bone-in pork rib chops", Description: "about 1 pound each, patted dry"},
		},
		{
			name:  "unit after comma",
			input: "2 veal shoulder arm or blade steaks, cut 1 inch thick",
			want:  model.Ingredient{Amount: 2, Unit: "", UnitCode: "", Name: "veal shoulder arm or blade steaks", Description: "cut 1 inch thick"},
		},
		{
			name:  "multiple units",
			input: "2 cups / 480 ml water",
			want:  model.Ingredient{Amount: 2, Unit: "cups", UnitCode: "cup", Name: "water"},
		},
		{
			name:  "merged unit",
			input: "500g carrots",
			want:  model.Ingredient{Amount: 500, Unit: "g", UnitCode: "g", Name: "carrots"},
		},
		{
			name:  "merged unit 2",
			input: "2EL Mayonnaise",
			want:  model.Ingredient{Amount: 2, Unit: "", Name: "EL Mayonnaise"},
		},
		{
			name:  "inches",
			input: "4 med. potatoes, pared & cut lengthwise, in 1/4\" slices",
			want:  model.Ingredient{Amount: 4, Unit: "", Name: "med. potatoes", Description: "pared & cut lengthwise, in 1/4\" slices"},
		},
		{
			name:  "slash fraction",
			input: "1/2 cup milk",
			want:  model.Ingredient{Amount: 0.5, Unit: "cup", UnitCode: "cup", Name: "milk"},
		},
		{
			name:  "unicode fraction slash",
			input: "1\u20442 cup milk",
			want:  model.Ingredient{Amount: 0.5, Unit: "cup", UnitCode: "cup", Name: "milk"},
		},
		{
			name:  "mixed number slash",
			input: "1 1/2 cups flour",
			want:  model.Ingredient{Amount: 1.5, Unit: "cups", UnitCode: "cup", Name: "flour"},
		},
		{
			name:  "decimal amount",
			input: "1.5 tablespoons olive oil",
			want:  model.Ingredient{Amount: 1.5, Unit: "tablespoons", UnitCode: "tbsp", Name: "olive oil"},
		},
		{
			name:  "word number",
			input: "two cups flour",
			want:  model.Ingredient{Amount: 2, Unit: "cups", UnitCode: "cup", Name: "flour"},
		},
		{
			name:  "article as one",
			input: "a pinch of salt",
			want:  model.Ingredient{Amount: 1, Unit: "pinch", UnitCode: "pinch", Name: "of salt"},
		},
		{
			name:  "word range",
			input: "1 to 2 cups sugar",
			want:  model.Ingredient{Amount: 1, MaxAmount: 2, Unit: "cups", UnitCode: "cup", Name: "sugar"},
		},
		{
			name:  "case insensitive unit",
			input: "2 CUPS water",
			want:  model.Ingredient{Amount: 2, Unit: "CUPS", UnitCode: "cup", Name: "water"},
		},
		{
			name:  "no amount",
			input: "salt",
			want:  model.Ingredient{Amount: 0, Unit: "", Name: "salt"},
		},
		{
			name:  "no unit",
			input: "3 eggs",
			want:  model.Ingredient{Amount: 3, Unit: "", Name: "eggs"},
		},
		{
			name:  "multiple commas",
			input: "flour, organic, finely sifted",
			want:  model.Ingredient{Amount: 0, Unit: "", Name: "flour", Description: "organic, finely sifted"},
		},
		{
			name:  "plus as separator",
			input: "salt + pepper to taste",
			want:  model.Ingredient{Amount: 0, Unit: "", Name: "salt + pepper to taste"},
		},
		{
			name:  "with size suffix",
			input: "2 large eggs",
			want:  model.Ingredient{Amount: 2, Unit: "", Name: "eggs", Description: "large"},
		},
		{
			name:  "with size suffix and comma",
			input: "3 medium potatoes, peeled and cut into 1/2\" slices",
			want:  model.Ingredient{Amount: 3, Unit: "", Name: "potatoes", Description: "medium, peeled and cut into 1/2\" slices"},
		},
		{
			name:  "unit after comma",
			input: "2 veal shoulder arm or blade steaks, cut 1 inch thick",
			want:  model.Ingredient{Amount: 2, Unit: "", UnitCode: "", Name: "veal shoulder arm or blade steaks", Description: "cut 1 inch thick"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseIngredient(tt.input, Options{Lang: "en"})
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkParseIngredient(b *testing.B) {
	for b.Loop() {
		_, _ = ParseIngredient("1 ½ cups diced tomatoes, divided", Options{Lang: "en"})
	}
}
