package testdata

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/utils"
)

var WebsiteExt = ".html"
var WebsiteJsonExt = ".json"
var RecipeExt = ".recipe.json"
var RecipeNewExt = ".recipe-new.json"
var WebsitesDir = currentPath() + "/websites/"

func currentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func AssertRecipe(t *testing.T, recipe *model.Recipe) {
	alias := utils.ParserAlias(recipe.Url)
	AssertRecipeAlias(t, recipe, alias)
}

func AssertRecipeAlias(t *testing.T, recipe *model.Recipe, alias string) {
	recipeData, err := json.MarshalIndent(recipe, "", "  ")
	assert.NoError(t, err)

	expectedRecipe, err := ioutil.ReadFile(WebsitesDir + alias + RecipeExt)
	assert.NoError(t, err)

	if !assert.Equal(t, string(expectedRecipe), string(recipeData)) {
		writeFileName := WebsitesDir + alias + RecipeNewExt
		if _, ok := os.LookupEnv("RECIPE_OVERRIDE"); ok {
			writeFileName = WebsitesDir + alias + RecipeExt
		}

		assert.NoError(t, ioutil.WriteFile(writeFileName, recipeData, 0644))
	}
}

func OptionallyMockRequests(t *testing.T) {
	if _, ok := os.LookupEnv("RECIPE_OFFLINE"); ok {
		MockRequests(t)
	}
}

func MockRequests(t *testing.T) {
	httpmock.Activate()

	httpmock.RegisterNoResponder(func(req *http.Request) (*http.Response, error) {
		data, err := mockResponse(req.URL.String(), req.Header.Get("Accept"))

		if err != nil {
			return httpmock.NewStringResponse(http.StatusInternalServerError, "HttpMock: "+err.Error()), nil
		} else {
			response := httpmock.NewBytesResponse(http.StatusOK, data)
			if req.Header.Get("Accept") == "application/json" {
				response.Header.Set("Content-Type", "application/json")
			} else {
				response.Header.Set("Content-Type", "text/html; charset=utf-8")
			}
			return response, nil
		}
	})

	t.Cleanup(httpmock.Deactivate)
}

func mockResponse(requestUrl string, accept string) ([]byte, error) {
	fileName := WebsitesDir + utils.ParserAlias(requestUrl)
	switch accept {
	case "", "text/html":
		fileName += WebsiteExt
	case "application/json":
		fileName += WebsiteJsonExt
	default:
		return nil, errors.New("unknown accept type")
	}

	return ioutil.ReadFile(fileName)
}
