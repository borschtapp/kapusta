package kapusta

import (
	"errors"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/parser"
	"borscht.app/kapusta/parser/opengraph"
	"borscht.app/kapusta/parser/schema"
	"borscht.app/kapusta/parser/website"
	"borscht.app/kapusta/utils"
)

// Scraper defines a function that fill a recipe from the input data
type Scraper = func(p *parser.InputData, r *model.Recipe) error

var scrapers = map[string]Scraper{
	"marleyspoon": website.ParseMarleySpoon,
}

func RegisterScraper(hostname string, fn func(p *parser.InputData, r *model.Recipe) error) {
	scrapers[hostname] = fn
}

func Scrape(input *parser.InputData) (*model.Recipe, error) {
	recipe := &model.Recipe{}
	recipe.Url = input.Url

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

	alias := utils.ParserAlias(input.Url)
	// fill recipe according to the alias parser implementation
	if aliasParser, ok := scrapers[alias]; ok {
		err := aliasParser(input, recipe)
		if err != nil {
			return nil, errors.New("alias parser error: " + err.Error())
		}
	}

	return recipe, nil
}

// ScrapeFile reads content and parses a recipe from the file
func ScrapeFile(fileName string) (*model.Recipe, error) {
	input, err := parser.FileInput(fileName, parser.Options{})
	if err != nil {
		return nil, err
	}

	return Scrape(input)
}

// ScrapeUrl retrieves and parses a recipe from the url
func ScrapeUrl(url string) (*model.Recipe, error) {
	input, err := parser.UrlInput(url, parser.Options{})
	if err != nil {
		return nil, err
	}

	return Scrape(input)
}
