package kapusta

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
