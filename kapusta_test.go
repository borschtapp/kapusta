package kapusta_test

import (
	"testing"

	"github.com/borschtapp/kapusta"
	"github.com/stretchr/testify/assert"
)

func TestParseIngredient(t *testing.T) {
	ing, err := kapusta.ParseIngredient("3 tablespoons olive oil")
	assert.NoError(t, err)
	assert.Equal(t, float64(3), ing.Amount)
	assert.Equal(t, "tablespoons", ing.Unit)
	assert.Equal(t, "olive oil", ing.Name)

	// German auto-detection
	ing, err = kapusta.ParseIngredient("1 EL Zucker")
	assert.NoError(t, err)
	assert.Equal(t, float64(1), ing.Amount)
	assert.Equal(t, "EL", ing.Unit)
	assert.Equal(t, "Zucker", ing.Name)
	assert.Equal(t, "tbsp", ing.UnitCode)

	// Spanish auto-detection
	ing, err = kapusta.ParseIngredient("2 dientes de ajo")
	assert.NoError(t, err)
	assert.Equal(t, float64(2), ing.Amount)
	assert.Equal(t, "dientes", ing.Unit)
	assert.Equal(t, "de ajo", ing.Name)
	assert.Equal(t, "clove", ing.UnitCode)

	// French auto-detection
	ing, err = kapusta.ParseIngredient("200g de farine")
	assert.NoError(t, err)
	assert.Equal(t, float64(200), ing.Amount)
	assert.Equal(t, "g", ing.Unit)
	assert.Equal(t, "de farine", ing.Name)
	assert.Equal(t, "g", ing.UnitCode)
}

func TestParseInstruction(t *testing.T) {
	ins, err := kapusta.ParseInstruction("Let it rest for 1 hour.")
	assert.NoError(t, err)
	assert.Len(t, ins.Timers, 1)
	assert.Equal(t, 3600, ins.Timers[0].Value)
	assert.Equal(t, `Let it rest for [1 hour](timer:3600).`, ins.Markdown)

	// German instruction with auto-detection
	ins, err = kapusta.ParseInstruction("Lass es 1 Stunde ruhen.")
	assert.NoError(t, err)
	assert.Len(t, ins.Timers, 1)
	assert.Equal(t, 3600, ins.Timers[0].Value)
	assert.Equal(t, `Lass es [1 Stunde](timer:3600) ruhen.`, ins.Markdown)
}
