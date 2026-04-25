package kapusta_test

import (
	"testing"

	"github.com/borschtapp/kapusta"
	"github.com/stretchr/testify/assert"
)

func TestParseIngredient(t *testing.T) {
	ing, err := kapusta.ParseIngredient("3 tablespoons olive oil", kapusta.IngredientOptions{Lang: "en"})
	assert.NoError(t, err)
	assert.Equal(t, float64(3), ing.Amount)
	assert.Equal(t, "tablespoons", ing.Unit)
	assert.Equal(t, "olive oil", ing.Name)
}

func TestParseInstruction(t *testing.T) {
	ins, err := kapusta.ParseInstruction("Let it rest for 1 hour.", kapusta.InstructionOptions{
		Lang: "en",
	})
	assert.NoError(t, err)
	assert.Len(t, ins.Timers, 1)
	assert.Equal(t, 3600, ins.Timers[0].Value)
	assert.Equal(t, `Let it rest for [1 hour](timer:3600).`, ins.Markdown)
}
