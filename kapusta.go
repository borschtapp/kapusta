package kapusta

import (
	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/ingredient"
	"github.com/borschtapp/kapusta/parser/instruction"
)

func ParseIngredient(str string, opts IngredientOptions) (model.Ingredient, error) {
	return ingredient.ParseIngredient(str, opts)
}

func ParseInstruction(str string, opts InstructionOptions) (model.Instruction, error) {
	return instruction.ParseInstruction(str, opts)
}

type IngredientOptions = ingredient.Options
type InstructionOptions = instruction.Options
type Ingredient = model.Ingredient
type Instruction = model.Instruction
type Timer = model.Timer
