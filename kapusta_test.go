package kapusta

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/scraper"
	"borscht.app/kapusta/testdata"
	"borscht.app/kapusta/utils"
)

func TestTestdataWebsites(t *testing.T) {
	testdata.MockRequests(t)

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
				input, err := scraper.ScrapeFile(path, model.Options{SkipText: true, SkipSchema: true})
				assert.NoError(t, err)
				assert.NotEmpty(t, input.Url)

				expected := utils.ParserAlias(input.Url)
				assert.NotRegexp(t, regexp.MustCompile(`^file://.+`), input.Url)
				assert.Regexp(t, regexp.MustCompile(`^`+expected+`\d*\.html$`), info.Name(), "Expected filename is "+expected+" for url "+input.Url)
			})
		}
		return nil
	})
}
