package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/scraper"
)

func TestParser(t *testing.T) {
	// recipeUrl := "http://allrecipes.com/recipe/231495/texas-boiled-beer-shrimp/"
	// recipeUrl := "https://www.kuchnia-domowa.pl/przepisy/dania-glowne/595-krokiety-z-pieczarkami-jajkiem-zoltym-serem"
	// recipeUrl := "https://www.allrecipes.com/recipe/231175/my-favorite-sloppy-joes/"
	recipeUrl := "https://www.hellofresh.com/recipes/uk-stir-fried-chinese-beef-5845b40b2e69d7259304d962"

	input, err := scraper.ScrapeUrl(recipeUrl, model.Options{})
	assert.Nil(t, err)

	recipe := &model.Recipe{}
	assert.Nil(t, Parse(input, recipe))

	assert.Equal(t, `{
  "name": "Rapid Stir-Fried Beef and Broccoli",
  "image": [
    "https://img.hellofresh.com/f_auto,fl_lossy,h_640,q_auto,w_1200/hellofresh_s3/image/uk-stir-friend-chinese-beef-b5fd1d10.jpg"
  ],
  "author": {
    "name": "HelloFresh"
  },
  "description": "One of the great things about Asian-style stir-fries is that they deliver maximum flavor in a snap. In this recipe, we’re tossing the classic combo of beef and broccoli with bouncy noodles and dressing them in a savory soy and hoisin-based sauce. It comes together so swiftly in the pan, it might even be easier than picking up the phone for takeout!",
  "totalTime": 20,
  "keywords": "Rapid, Spicy",
  "recipeYield": 2,
  "recipeCategory": "main course",
  "recipeCuisine": "Asian",
  "nutrition": {
    "calories": 754,
    "carbohydrateContent": 77,
    "cholesterolContent": 102,
    "fatContent": 29,
    "fiberContent": 6,
    "proteinContent": 54,
    "saturatedFatContent": 6,
    "sodiumContent": 1620,
    "sugarContent": 9
  },
  "recipeIngredient": [
    "12 ounce Beef Sirloin Tips",
    "2 unit Scallions",
    "2 clove Garlic",
    "1 tablespoon Cornstarch",
    "1 thumb Ginger",
    "16 ounce Yakisoba Noodles",
    "1 unit Ketchup",
    "4 unit Soy Sauce",
    "1 jar Hoisin Sauce Jar",
    "8 ounce Broccoli Florets",
    "1 tablespoon Sesame Oil",
    "1 teaspoon Sriracha",
    "4 teaspoon Vegetable Oil",
    "unit Salt",
    "unit Pepper"
  ],
  "recipeInstructions": [
    {
      "text": "Wash and dry all produce. Bring a large pot of salted water to a boil. Trim and thinly slice scallions. Mince or grate garlic. Peel and mince ginger. Whisk together sesame oil, 1 TBSP ketchup, soy sauce, 1½ TBSP hoisin sauce, and 1 TBSP water in a small bowl."
    },
    {
      "text": "Add broccoli to boiling water and cook until tender but still crisp, 3-4 minutes. Drain and rinse under cold water. Set aside."
    },
    {
      "text": "Toss steak tips with cornstarch in a large bowl. Season generously with salt and pepper. Heat a large drizzle of oil in a large pan over high heat. (TIP: If you have a nonstick pan, break it out.) Toss in steak tips and cook to desired doneness, 3-4 minutes. Remove and set aside."
    },
    {
      "text": "Heat a drizzle of oil in same pan over medium heat. Add garlic, ginger, and scallions and cook until fragrant, 1 minute, tossing. Toss in half the noodles from the package (we sent more) and a drizzle of oil. Break up noodles until they no longer stick together, using tongs or two wooden spoons."
    },
    {
      "text": "Pour in 1 cup water, cover, and steam until noodles are tender, 3 minutes. (TIP: If your pan doesn’t have a lid, carefully cover it with aluminum foil.) Uncover, increase heat to medium-high, and toss until noodles are tender, 3-4 minutes. Add sauce and toss to coat. Cook until sauce is thickened, 1 minute."
    },
    {
      "text": "Toss broccoli and steak into noodles to warm through. Season with as much sriracha as you like (careful, it’s spicy). Season with salt and pepper. Divide between plates and serve."
    }
  ],
  "datePublished": "2016-12-05T18:38:03Z"
}`, recipe.String())
}
