package website

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/parser"
	"borscht.app/kapusta/scraper"
	"borscht.app/kapusta/testdata"
)

func TestMarleySpoon(t *testing.T) {
	testdata.OptionallyMockRequests(t)

	input, err := scraper.ScrapeUrl("https://marleyspoon.de/menu/113813-glasierte-veggie-burger-mit-roestkartoffeln-und-apfel-gurken-salat", model.Options{})
	assert.NoError(t, err)

	recipe, err := parser.GeneralParsers(input)
	assert.NoError(t, err)

	err = ParseMarleySpoon(input, recipe)
	assert.NoError(t, err)

	testdata.AssertRecipe(t, recipe)
}
