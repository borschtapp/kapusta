package kapusta

import (
	"github.com/borschtapp/kapusta/parser/ingredient"
	"github.com/borschtapp/krip/model"
)

func ParseIngredient(str string, lang string) (*model.Ingredient, error) {
	return ingredient.Parse(str, lang)
}
