package kapusta

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/publicsuffix"

	"borscht.app/kapusta/parser"
)

var inExt = ".html"
var outExt = ".recipe.json"
var dir = "testdata/websites/"

func TestWebsites(t *testing.T) {
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		assert.NoError(t, err)

		if !info.IsDir() && strings.HasSuffix(info.Name(), inExt) {
			t.Run(info.Name(), func(t *testing.T) {
				recipe, err := ScrapeFile(path)
				assert.NoError(t, err)

				jsonData, err := json.MarshalIndent(recipe, "", "  ")
				assert.NoError(t, err)
				err = ioutil.WriteFile(strings.Replace(path, inExt, outExt, 1), jsonData, 0644)
				assert.NoError(t, err)
			})
		}
		return nil
	})
}

func TestTestdataFilenames(t *testing.T) {
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		assert.NoError(t, err)

		if strings.HasSuffix(info.Name(), inExt) {
			t.Run(info.Name(), func(t *testing.T) {
				input, err := parser.FileInput(path, parser.Options{SkipText: true, SkipSchema: true})
				assert.NoError(t, err)
				recipe, err := Scrape(input)
				assert.NoError(t, err)

				assert.NotEmpty(t, recipe.Url)
				expected := getFilenameForUrl(recipe.Url)
				assert.NotRegexp(t, regexp.MustCompile(`^file://.+`), recipe.Url)
				assert.Regexp(t, regexp.MustCompile(`^`+expected+`\d*\.html$`), info.Name(), "Expected filename is "+expected+" for url "+recipe.Url)
			})
		}
		return nil
	})
}

func getFilenameForUrl(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	file := strings.ToLower(u.Hostname())
	file = strings.Replace(file, "www.", "", 1)

	suffix, _ := publicsuffix.PublicSuffix(file)
	file = strings.Replace(file, "."+suffix, "", 1)
	file = strings.ReplaceAll(file, ".", "_")
	return file
}
