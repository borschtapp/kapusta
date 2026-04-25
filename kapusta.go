package kapusta

import (
	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/ingredient"
)

func ParseIngredient(str string, opts IngredientOptions) (model.Ingredient, error) {
	return ingredient.Parse(str, opts)
}

type IngredientOptions = ingredient.Options
type Ingredient = model.Ingredient
