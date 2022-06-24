package website

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/parser"
	"borscht.app/kapusta/scraper"
	"borscht.app/kapusta/testdata"
)

func TestFitMenCook(t *testing.T) {
	testdata.OptionallyMockRequests(t)

	input, err := scraper.ScrapeUrl("https://fitmencook.com/healthy-chili-recipe/", model.Options{})
	assert.NoError(t, err)

	recipe, err := parser.GeneralParsers(input)
	assert.NoError(t, err)

	err = ParseFitMenCook(input, recipe)
	assert.NoError(t, err)

	testdata.AssertRecipe(t, recipe)
}
