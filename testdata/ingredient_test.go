package testdata

import (
	"bufio"
	"github.com/borschtapp/kapusta"
	"github.com/borschtapp/kapusta/model"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestParseTestdataIngredients(t *testing.T) {
	t.Parallel()

	file, err := os.Open(TestdataDir + "ingredients.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()

		var ingredients []*model.Ingredient
		t.Run(line, func(t *testing.T) {
			ing, err := kapusta.ParseIngredient(line, "en")
			assert.NoError(t, err)
			ingredients = append(ingredients, ing)
			println(ing.String())
		})
	}
}
