package instruction

import (
	"fmt"
	"testing"

	"github.com/borschtapp/kapusta/model"
	"github.com/stretchr/testify/assert"
)

func ExampleParseInstruction() {
	inst, _ := ParseInstruction("Bake potatoes for 30 minutes.", Options{
		Lang:        "en",
		Ingredients: []model.IngredientRef{{Name: "potatoes"}},
	})

	fmt.Println("Text:", inst.Text)
	fmt.Println("Markdown:", inst.Markdown)
	fmt.Printf("Timer Value: %d, Raw: %s\n", inst.Timers[0].Value, inst.Timers[0].Raw)
	fmt.Printf("Ingredient ID: %s, Name: %s\n", inst.Ingredients[0].ID, inst.Ingredients[0].Name)

	// Output:
	// Text: Bake potatoes for 30 minutes.
	// Markdown: Bake [potatoes](ingredient:kap0) for [30 minutes](timer:1800).
	// Timer Value: 1800, Raw: 30 minutes
	// Ingredient ID: kap0, Name: potatoes
}

func TestParseInstruction(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		opts      Options
		want      model.Instruction
		expectErr bool
	}{
		{
			name:  "timer",
			input: "Bake for 15 mins.",
			opts:  Options{Lang: "en"},
			want: model.Instruction{
				Text:     "Bake for 15 mins.",
				Markdown: `Bake for [15 mins](timer:900).`,
				Timers:   []model.Timer{{Value: 900, Raw: "15 mins"}},
			},
		},
		{
			name:  "timer seconds",
			input: "Stir for 30 seconds.",
			opts:  Options{Lang: "en"},
			want: model.Instruction{
				Text:     "Stir for 30 seconds.",
				Markdown: `Stir for [30 seconds](timer:30).`,
				Timers:   []model.Timer{{Value: 30, Raw: "30 seconds"}},
			},
		},
		{
			name:  "timer range",
			input: "Cook it for 2-3 minutes.",
			opts:  Options{Lang: "en"},
			want: model.Instruction{
				Text:     "Cook it for 2-3 minutes.",
				Markdown: `Cook it for [2-3 minutes](timer:120?max=180).`,
				Timers:   []model.Timer{{Value: 120, MaxValue: 180, Raw: "2-3 minutes"}},
			},
		},
		{
			name:  "ingredients",
			input: "Add salt and pepper.",
			opts:  Options{Lang: "en", Ingredients: []model.IngredientRef{{Name: "salt"}, {Name: "pepper"}}},
			want: model.Instruction{
				Text:        "Add salt and pepper.",
				Markdown:    "Add [salt](ingredient:kap0) and [pepper](ingredient:kap1).",
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "salt"}, {ID: "kap1", Name: "pepper"}},
			},
		},
		{
			name:  "ingredients from units",
			input: "Add 1 cup flour and 2 tbsp sugar.",
			opts:  Options{Lang: "en"},
			want: model.Instruction{
				Text:        "Add 1 cup flour and 2 tbsp sugar.",
				Markdown:    "Add [1 cup](amount:1?unit=cup) [flour](ingredient:kap0) and [2 tbsp](amount:2?unit=tbsp) [sugar](ingredient:kap1).",
				Amounts:     []model.Amount{{Value: 1, Unit: "cup", Raw: "1 cup"}, {Value: 2, Unit: "tbsp", Raw: "2 tbsp"}},
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "flour"}, {ID: "kap1", Name: "sugar"}},
			},
		},
		{
			name:  "timer and ingredient",
			input: "Bake potatoes for 30 minutes.",
			opts:  Options{Lang: "en", Ingredients: []model.IngredientRef{{Name: "potatoes"}}},
			want: model.Instruction{
				Text:        "Bake potatoes for 30 minutes.",
				Markdown:    `Bake [potatoes](ingredient:kap0) for [30 minutes](timer:1800).`,
				Timers:      []model.Timer{{Value: 1800, Raw: "30 minutes"}},
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "potatoes"}},
			},
		},
		{
			name:      "invalid language",
			input:     "Bake for 15 mins.",
			opts:      Options{Lang: "xx"},
			expectErr: true,
		},
		{
			name:  "multiple timers",
			input: "Bake for 30 minutes, then rest for 10 minutes.",
			opts:  Options{Lang: "en"},
			want: model.Instruction{
				Text:     "Bake for 30 minutes, then rest for 10 minutes.",
				Markdown: `Bake for [30 minutes](timer:1800), then rest for [10 minutes](timer:600).`,
				Timers:   []model.Timer{{Value: 1800, Raw: "30 minutes"}, {Value: 600, Raw: "10 minutes"}},
			},
		},
		{
			name:  "multi word ingredient",
			input: "Add tomato paste and a tomato.",
			opts:  Options{Lang: "en", Ingredients: []model.IngredientRef{{Name: "tomato"}, {Name: "tomato paste"}}},
			want: model.Instruction{
				Text:        "Add tomato paste and a tomato.",
				Markdown:    "Add [tomato paste](ingredient:kap0) and a [tomato](ingredient:kap1).",
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "tomato paste"}, {ID: "kap1", Name: "tomato"}},
			},
		},
		{
			name:  "temperature and amounts",
			input: "Heat 500 ml water to 100°C.",
			opts:  Options{Lang: "en"},
			want: model.Instruction{
				Text:     "Heat 500 ml water to 100°C.",
				Markdown: `Heat [500 ml](amount:500?unit=ml) [water](ingredient:kap0) to [100°C](temp:100?unit=C).`,
				Amounts:  []model.Amount{{Value: 500, Unit: "ml", Raw: "500 ml"}},
				Temperatures: []model.Temperature{
					{Value: 100, Unit: "C", Raw: "100°C"},
				},
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "water"}},
			},
		},
		{
			name:  "custom IDPrefix",
			input: "Add 1 cup flour.",
			opts:  Options{Lang: "en", IDPrefix: "ing"},
			want: model.Instruction{
				Text:        "Add 1 cup flour.",
				Markdown:    "Add [1 cup](amount:1?unit=cup) [flour](ingredient:ing0).",
				Amounts:     []model.Amount{{Value: 1, Unit: "cup", Raw: "1 cup"}},
				Ingredients: []model.IngredientRef{{ID: "ing0", Name: "flour"}},
			},
		},
		{
			name:  "overlap: temp wins over amount (ambiguous C)",
			input: "Heat to 100 C",
			opts:  Options{Lang: "en"},
			want: model.Instruction{
				Text:     "Heat to 100 C",
				Markdown: "Heat to [100 C](temp:100?unit=C)",
				Temperatures: []model.Temperature{
					{Value: 100, Unit: "C", Raw: "100 C"},
				},
				Amounts: nil, // Without P1, this would have contained "100 cups"
			},
		},
		{
			name:  "overlap: amount wins over inferred ingredient",
			input: "Add 1 cup salt.",
			opts:  Options{Lang: "en"},
			want: model.Instruction{
				Text:     "Add 1 cup salt.",
				Markdown: "Add [1 cup](amount:1?unit=cup) [salt](ingredient:kap0).",
				Amounts: []model.Amount{
					{Value: 1, Unit: "cup", Raw: "1 cup"},
				},
				Ingredients: []model.IngredientRef{
					{ID: "kap0", Name: "salt"},
				},
			},
		},
		{
			name:  "case insensitive ingredient matching",
			input: "Add Salt.",
			opts:  Options{Lang: "en", Ingredients: []model.IngredientRef{{Name: "salt"}}},
			want: model.Instruction{
				Text:        "Add Salt.",
				Markdown:    "Add [Salt](ingredient:kap0).",
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "salt"}},
			},
		},
		{
			name:  "longest match wins for ingredients",
			input: "Add olive oil.",
			opts:  Options{Lang: "en", Ingredients: []model.IngredientRef{{Name: "oil"}, {Name: "olive oil"}}},
			want: model.Instruction{
				Text:        "Add olive oil.",
				Markdown:    "Add [olive oil](ingredient:kap0).",
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "olive oil"}},
			},
		},
		{
			name:  "preserve existing ingredient IDs",
			input: "Add salt.",
			opts:  Options{Lang: "en", Ingredients: []model.IngredientRef{{ID: "custom-id", Name: "salt"}}},
			want: model.Instruction{
				Text:        "Add salt.",
				Markdown:    "Add [salt](ingredient:custom-id).",
				Ingredients: []model.IngredientRef{{ID: "custom-id", Name: "salt"}},
			},
		},
		{
			name:  "repeated ingredient uses same ID",
			input: "Add 1 cup water and then another cup water.",
			opts:  Options{Lang: "en", Ingredients: []model.IngredientRef{{Name: "water"}}},
			want: model.Instruction{
				Text:        "Add 1 cup water and then another cup water.",
				Markdown:    "Add [1 cup](amount:1?unit=cup) [water](ingredient:kap0) and then another cup [water](ingredient:kap0).",
				Amounts:     []model.Amount{{Value: 1, Unit: "cup", Raw: "1 cup"}},
				Ingredients: []model.IngredientRef{{ID: "kap0", Name: "water"}},
			},
		},
		{
			name:  "inferred multi-word ingredient (known limitation check)",
			input: "Add 1 cup tomato paste.",
			opts:  Options{Lang: "en"},
			want: model.Instruction{
				Text:     "Add 1 cup tomato paste.",
				Markdown: "Add [1 cup](amount:1?unit=cup) [tomato](ingredient:kap0) paste.",
				Amounts:  []model.Amount{{Value: 1, Unit: "cup", Raw: "1 cup"}},
				Ingredients: []model.IngredientRef{
					{ID: "kap0", Name: "tomato"},
				},
			},
		},
		{
			name:  "multi-line instruction",
			input: "Bake for 30 minutes.\nThen rest for 10 minutes.",
			opts:  Options{Lang: "en"},
			want: model.Instruction{
				Text:     "Bake for 30 minutes.\nThen rest for 10 minutes.",
				Markdown: "Bake for [30 minutes](timer:1800).\nThen rest for [10 minutes](timer:600).",
				Timers: []model.Timer{
					{Value: 1800, Raw: "30 minutes"},
					{Value: 600, Raw: "10 minutes"},
				},
			},
		},
	}

	normalize := func(i model.Instruction) model.Instruction {
		if i.Timers == nil {
			i.Timers = []model.Timer{}
		}
		if i.Ingredients == nil {
			i.Ingredients = []model.IngredientRef{}
		}
		return i
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInstruction(tt.input, tt.opts)
			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, normalize(tt.want), normalize(got))
		})
	}
}

func BenchmarkParseInstruction(b *testing.B) {
	for b.Loop() {
		_, _ = ParseInstruction("Bake potatoes for 30 minutes at 200 degrees.", Options{Lang: "en"})
	}
}
