package dictionary_test

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/ingredient"
	"github.com/borschtapp/kapusta/parser/instruction"
	"github.com/stretchr/testify/assert"
)

type ingredientTestCase struct {
	give string
	want model.Ingredient
}

func runIngredientTests(t *testing.T, lang string, tests []ingredientTestCase) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := ingredient.ParseIngredient(tt.give, ingredient.Options{Lang: lang})
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

type instructionTestCase struct {
	give        string
	ingredients []string
	want        model.Instruction
}

func normalizeInstruction(i model.Instruction) model.Instruction {
	if i.Timers == nil {
		i.Timers = []model.Timer{}
	}
	if i.Ingredients == nil {
		i.Ingredients = []model.IngredientRef{}
	}
	if i.Temperatures == nil {
		i.Temperatures = []model.Temperature{}
	}
	if i.Amounts == nil {
		i.Amounts = []model.Amount{}
	}
	return i
}

func runInstructionTests(t *testing.T, lang string, tests []instructionTestCase) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			var ingredRefs []model.IngredientRef
			for _, name := range tt.ingredients {
				ingredRefs = append(ingredRefs, model.IngredientRef{Name: name})
			}

			got, err := instruction.ParseInstruction(tt.give, instruction.Options{
				Lang:        lang,
				Ingredients: ingredRefs,
			})
			assert.NoError(t, err)
			assert.Equal(t, normalizeInstruction(tt.want), normalizeInstruction(got))
		})
	}
}
