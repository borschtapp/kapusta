package nlp

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/scraper"
)

func TestParser(t *testing.T) {
	recipeUrl := "https://www.hellofresh.com/recipes/uk-stir-fried-chinese-beef-5845b40b2e69d7259304d962"

	input, err := scraper.ScrapeUrl(recipeUrl, model.Options{})
	assert.Nil(t, err)

	recipe := &model.Recipe{}
	assert.Nil(t, Parse(input, recipe))

	assert.Equal(t, `{
  "recipeIngredient": [
    " 2 clove ",
    " sugar 9 g ",
    " wash and dry all produce bring a large pot of salted water to a boil trim and thinly slice scallions mince or grate garlic peel and mince ginger whisk together sesame oil 1 tbsp ketchup soy sauce 1½ tbsp hoisin sauce and 1 tbsp water in a small bowl ",
    " img src https //img hellofresh com/f auto fl lossy q auto w 385/hellofresh s3/step/5845b40b2e69d7259304d962 1 0 d30146dd jpg alt cook broccoli class dsel dsed / ",
    " img src https //img hellofresh com/f auto fl lossy q auto w 385/hellofresh s3/step/5845b40b2e69d7259304d962 1 0 d30146dd jpg alt cook broccoli class dsel dsed / ",
    " add broccoli to boiling water and cook until tender but still crisp 3 4 minutes drain and rinse under cold water set aside ",
    " img src https //img hellofresh com/f auto fl lossy q auto w 385/hellofresh s3/step/5845b40b2e69d7259304d962 2 0 94993b41 jpg alt cook beef class dsel dsed / ",
    " img src https //img hellofresh com/f auto fl lossy q auto w 385/hellofresh s3/step/5845b40b2e69d7259304d962 2 0 94993b41 jpg alt cook beef class dsel dsed / ",
    " toss steak tips with cornstarch in a large bowl season generously with salt and pepper heat a large drizzle of oil in a large pan over high heat toss in steak tips and cook to desired doneness 3 4 minutes remove and set aside "
  ],
  "recipeInstructions": [
    {
      "text": "Toss steak tips with cornstarch in a large bowl. Season generously with salt and pepper. Heat a large drizzle of oil in a large pan over high heat. (TIP: If you have a nonstick pan, break it out.) Toss in steak tips and cook to desired doneness, 3-4 minutes. Remove and set aside."
    },
    {
      "text": "u003cimg src=\"https://img.hellofresh.com/f_auto,fl_lossy,q_auto,w_385/hellofresh_s3/step/5845b40b2e69d7259304d962-3-0-32afb5cb.jpg\" alt=\"Cook aromatics and noodles\" class=\"dsel dsed\"/>"
    },
    {
      "text": "<img src=\"https://img.hellofresh.com/f_auto,fl_lossy,q_auto,w_385/hellofresh_s3/step/5845b40b2e69d7259304d962-3-0-32afb5cb.jpg\" alt=\"Cook aromatics and noodles\" class=\"dsel dsed\"/>"
    },
    {
      "text": "Heat a drizzle of oil in same pan over medium heat. Add garlic, ginger, and scallions and cook until fragrant, 1 minute, tossing. Toss in half the noodles from the package (we sent more) and a drizzle of oil. Break up noodles until they no longer stick together, using tongs or two wooden spoons."
    },
    {
      "text": "<img src=\"https://img.hellofresh.com/f_auto,fl_lossy,q_auto,w_385/hellofresh_s3/step/5845b40b2e69d7259304d962-4-0-ed545ece.jpg\" alt=\"Steam noodles\" class=\"dsel dsed\"/>"
    },
    {
      "text": "<img src=\"https://img.hellofresh.com/f_auto,fl_lossy,q_auto,w_385/hellofresh_s3/step/5845b40b2e69d7259304d962-4-0-ed545ece.jpg\" alt=\"Steam noodles\" class=\"dsel dsed\"/>"
    },
    {
      "text": "Pour in 1 cup water, cover, and steam until noodles are tender, 3 minutes. (TIP: If your pan doesn’t have a lid, carefully cover it with aluminum foil.) Uncover, increase heat to medium-high, and toss until noodles are tender, 3-4 minutes. Add sauce and toss to coat. Cook until sauce is thickened, 1 minute."
    },
    {
      "text": "<img src=\"https://img.hellofresh.com/f_auto,fl_lossy,q_auto,w_385/hellofresh_s3/step/5845b40b2e69d7259304d962-5-0-a7b12cbb.jpg\" alt=\"Finish and service\" class=\"dsel dsed\"/>"
    },
    {
      "text": "<img src=\"https://img.hellofresh.com/f_auto,fl_lossy,q_auto,w_385/hellofresh_s3/step/5845b40b2e69d7259304d962-5-0-a7b12cbb.jpg\" alt=\"Finish and service\" class=\"dsel dsed\"/>"
    },
    {
      "text": "Toss broccoli and steak into noodles to warm through. Season with as much sriracha as you like (careful, it’s spicy). Season with salt and pepper. Divide between plates and serve."
    }
  ]
}`, recipe.String())
}
