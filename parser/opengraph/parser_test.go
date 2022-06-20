package opengraph

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/parser"
)

func TestParser(t *testing.T) {
	// recipeUrl := "http://allrecipes.com/recipe/231495/texas-boiled-beer-shrimp/"
	// recipeUrl := "https://www.kuchnia-domowa.pl/przepisy/dania-glowne/595-krokiety-z-pieczarkami-jajkiem-zoltym-serem"
	// recipeUrl := "https://www.allrecipes.com/recipe/231175/my-favorite-sloppy-joes/"
	recipeUrl := "https://www.hellofresh.com/recipes/uk-stir-fried-chinese-beef-5845b40b2e69d7259304d962"

	input, err := parser.UrlInput(recipeUrl, parser.Options{})
	assert.Nil(t, err)

	recipe := &model.Recipe{}
	assert.Nil(t, Parse(input, recipe))

	assert.Equal(t, `{
  "name": "Rapid Stir-Fried Beef Recipe | HelloFresh",
  "image": [
    "https://img.hellofresh.com/f_auto,fl_lossy,h_640,q_auto,w_1200/hellofresh_s3/image/uk-stir-friend-chinese-beef-b5fd1d10.jpg"
  ],
  "description": "One of the great things about Asian-style stir-fries is that they deliver maximum flavor in a snap. In this recipe, we’re tossing the classic combo of beef and broccoli with bouncy noodles and dressing them in a savory soy and hoisin-based sauce. It comes together so swiftly in the pan, it might even be easier than picking up the phone for takeout!",
  "url": "https://www.hellofresh.com/recipes/uk-stir-fried-chinese-beef-5845b40b2e69d7259304d962",
  "siteName": "HelloFresh",
  "language": "en_US"
}`, recipe.String())
}
