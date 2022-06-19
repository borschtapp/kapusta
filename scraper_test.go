package kapusta

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"borscht.app/kapusta/parser"
	"borscht.app/kapusta/testdata"
	"borscht.app/kapusta/utils"
)

func TestAllWebsites(t *testing.T) {
	_ = filepath.Walk(testdata.WebsitesDir, func(path string, info os.FileInfo, err error) error {
		assert.NoError(t, err)

		if !info.IsDir() && strings.HasSuffix(info.Name(), testdata.WebsiteExt) {
			t.Run(info.Name(), func(t *testing.T) {
				recipe, err := ScrapeFile(path)
				assert.NoError(t, err)

				testdata.AssertRecipeAlias(t, recipe, strings.TrimSuffix(info.Name(), testdata.WebsiteExt))
			})
		}
		return nil
	})
}

func TestTestdataFilenames(t *testing.T) {
	_ = filepath.Walk(testdata.WebsitesDir, func(path string, info os.FileInfo, err error) error {
		assert.NoError(t, err)

		if strings.HasSuffix(info.Name(), testdata.WebsiteExt) {
			t.Run(info.Name(), func(t *testing.T) {
				input, err := parser.FileInput(path, parser.Options{SkipText: true, SkipSchema: true})
				assert.NoError(t, err)
				recipe, err := Scrape(input)
				assert.NoError(t, err)

				assert.NotEmpty(t, recipe.Url)
				expected := utils.ParserAlias(recipe.Url)
				assert.NotRegexp(t, regexp.MustCompile(`^file://.+`), recipe.Url)
				assert.Regexp(t, regexp.MustCompile(`^`+expected+`\d*\.html$`), info.Name(), "Expected filename is "+expected+" for url "+recipe.Url)
			})
		}
		return nil
	})
}
