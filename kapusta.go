package kapusta

import (
	"errors"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/parser"
	"borscht.app/kapusta/parser/website"
	"borscht.app/kapusta/scraper"
	"borscht.app/kapusta/utils"
)

// Parser defines a function that fill a recipe from the input data
type Parser = func(p *model.InputData, r *model.Recipe) error

var parsers = map[string]Parser{
	"marleyspoon": website.ParseMarleySpoon,
	"fitmencook":  website.ParseFitMenCook,
}

func RegisterParser(hostname string, fn Parser) {
	parsers[hostname] = fn
}

func Parse(input *model.InputData) (*model.Recipe, error) {
	recipe, err := parser.GeneralParsers(input)
	if err != nil {
		return nil, err
	}

	alias := utils.ParserAlias(input.Url)
	// fill recipe according to the alias parser implementation
	if aliasParser, ok := parsers[alias]; ok {
		err := aliasParser(input, recipe)
		if err != nil {
			return nil, errors.New("alias parser error: " + err.Error())
		}
	}

	return recipe, nil
}

// ScrapeFile reads content and parses a recipe from the file
func ScrapeFile(fileName string) (*model.Recipe, error) {
	input, err := scraper.ScrapeFile(fileName, model.Options{})
	if err != nil {
		return nil, err
	}

	return Parse(input)
}

// ScrapeUrl retrieves and parses a recipe from the url
func ScrapeUrl(url string) (*model.Recipe, error) {
	input, err := scraper.ScrapeUrl(url, model.Options{})
	if err != nil {
		return nil, err
	}

	return Parse(input)
}
