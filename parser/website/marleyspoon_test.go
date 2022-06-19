package website

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/parser"
	"borscht.app/kapusta/parser/opengraph"
	"borscht.app/kapusta/parser/schema"
	"borscht.app/kapusta/testdata"
)

func TestMarleySpoon(t *testing.T) {
	testdata.MockRequests(t)

	input, err := parser.UrlInput("https://marleyspoon.de/menu/113813-glasierte-veggie-burger-mit-roestkartoffeln-und-apfel-gurken-salat", parser.Options{})
	assert.NoError(t, err)

	recipe := &model.Recipe{Url: input.Url}
	err = opengraph.Parse(input, recipe)
	assert.NoError(t, err)
	err = schema.Parse(input, recipe)
	assert.NoError(t, err)
	err = ParseMarleySpoon(input, recipe)
	assert.NoError(t, err)
	testdata.AssertRecipe(t, recipe)
}
