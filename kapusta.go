package kapusta

import (
	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/ingredient"
)

func ParseIngredient(str string, lang string) (*model.Ingredient, error) {
	return ingredient.Parse(str, lang)
}

type Ingredient = model.Ingredient
