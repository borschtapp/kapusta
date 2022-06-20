package parser

import (
	"errors"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/parser/opengraph"
	"borscht.app/kapusta/parser/schema"
)

func GeneralParsers(input *model.InputData) (*model.Recipe, error) {
	recipe := &model.Recipe{Url: input.Url}

	// fill recipe with OpenGraph metadata
	err := opengraph.Parse(input, recipe)
	if err != nil {
		return nil, errors.New("opengraph error: " + err.Error())
	}

	// fill recipe with schema.org/Recipe metadata
	err = schema.Parse(input, recipe)
	if err != nil {
		return nil, errors.New("schema error: " + err.Error())
	}

	return recipe, nil
}
