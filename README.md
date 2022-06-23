# Kapusta (Brassica oleracea)

A Go library for scraping culinary recipes from any website or file.

---

This library is partially a Go port of [recipe-scrapers](https://github.com/hhursev/recipe-scrapers/),
but with many improvements that I found useful.

When I started looking for cooking planning apps, I found that they probably use a common scraper that
lacks some functionality. They all have a common flaws, and I wasn't satisfied, so I decided to create my own.

_Note:_ Of course, this is WIP, and I'm still learning how to use Go. Found it fun, but difficult to switch after OOP.

## Install

```
go get borscht.app/kapusta
```

## Features
- [x] The `Recipe` struct (object) returned by the library is defined by https://schema.org/Recipe
- [x] Removes empty, duplicate values on the fly
- [x] Fast, extremely fast

### TODO list
- [ ] add more website parsers

## Usage

### Scrape recipe from web

Firstly, you can parse the ingredients and directions from any website. For instance here's one I found on the [Chowdown](https://chowdown.io/recipes/mushroom-risotto.html):

```go
recipe, _ := kapusta.ScrapeUrl("https://chowdown.io/recipes/mushroom-risotto.html")
fmt.Println(recipe)
```
```json
{
  "name": "Mushroom Risotto",
  "image": [
    "https://chowdown.io/images/mushroom-risotto.jpg"
  ],
  "description": "Risotto is one of those recipes that sound tough but really isn’t that tricky (but is totally tasty). This recipe is a mash between a Jamie Oliver method and a Tasty 101 video. Jamie asked for dried mushrooms (proved to be hard to find and expensive here) and Tasty didn’t have enough mushroom punch (so we added back the roasted topper). The best of both worlds, and easier than I’ve previously imagined!",
  "recipeIngredient": [
    "10oz shitake mushrooms (we usually buy two 5oz grocery packages)",
    "1 cup arborio rice",
    "2 cups chicken stock",
    "4 tbsp butter (halved)",
    "1 small onion, diced",
    "6 cloves garlic",
    "1.5 cups white wine",
    "1 cup beer (drink the rest)",
    "half a lemon",
    "1 cup fresh grated parmesan cheese, plus more for serving",
    "sprig of parsley (or your favorite herb)",
    "salt",
    "pepper"
  ],
  "recipeInstructions": [
    {
      "text": "Add the chicken stock, 1 cup white wine, and 1 cup beer to a stock pock and bring to a simmer (4 cups total). Reduce heat or set aside."
    },
    {
      "text": "Add 2 tbsp butter, garlic, and onions to a deep pan and cook until translucent (a couple minutes). Add one package (or half) of your mushrooms and cook down. Salt and pepper to taste (small dash of salt, hearty crack of pepper for me, please)."
    },
    {
      "text": "Add the rice into the pan and toss to coat. Toast for a few minutes."
    },
    {
      "text": "Add the remaining white wine and juice from half a lemon to deglaze the pan, scraping up all the good bits."
    },
    {
      "text": "Add a large ladle (or two) of the hot stock mixture into the rice pan. Here’s where we get patient. Cook the stock down (near dry) before adding another ladle. Repeat."
    },
    {
      "text": "Continue to add a ladle at a time, letting it fully absord while stirring, until the stock is gone or the rice is tender. Fair game to taste the rice a few times in search of the perfect tenderness."
    },
    {
      "text": "Meanwhile, let’s roast our other pack of fresh mushrooms. Get the remaining half on a sheet pan and under a 450° broiler. Keep an eye on them but cook til crispy and brown on the edges."
    },
    {
      "text": "When the rice is tender and the mushrooms in the oven are crisp, stir your parmesan and remaining 2tbps butter into the rice. Plate your risotto and top with our roasted mushrooms, our herb, more parmesan, and a crack of black pepper."
    }
  ],
  "url": "https://chowdown.io/recipes/mushroom-risotto.html"
}

```
