package kapusta

import (
	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/ingredient"
	"github.com/borschtapp/kapusta/parser/instruction"
	"github.com/borschtapp/kapusta/parser/util"
)

func ParseIngredient(str string, opts ...IngredientOptions) (model.Ingredient, error) {
	return ingredient.ParseIngredient(str, util.First(opts))
}

func ParseInstruction(str string, opts ...InstructionOptions) (model.Instruction, error) {
	return instruction.ParseInstruction(str, util.First(opts))
}

type IngredientOptions = ingredient.Options
type InstructionOptions = instruction.Options
type Ingredient = model.Ingredient
type IngredientRef = model.IngredientRef
type Instruction = model.Instruction
type Timer = model.Timer
type Temperature = model.Temperature
type Amount = model.Amount
