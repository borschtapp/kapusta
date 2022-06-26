package nlp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIngredientsInString(t *testing.T) {
	line := SanitizeLine("3 1/2 cup chilled oil or lemony vinegar with lime (vegetable or canola oil)")
	wpi := GetIngredientsInString(line)
	assert.Equal(t, "oil", wpi[0].Word)
	assert.Equal(t, 3, len(wpi))
	assert.Equal(t, 17, wpi[0].Position)
	assert.Equal(t, 31, wpi[1].Position)
	assert.Equal(t, 44, wpi[2].Position)
	fmt.Println(wpi)

	wp := GetNumbersInString(line)
	assert.Equal(t, 2, len(wp))
	assert.Equal(t, "3", wp[0].Word)
	assert.Equal(t, "½", wp[1].Word)

	wpm := GetMeasuresInString(line)
	assert.Equal(t, 1, len(wpm))
	assert.Equal(t, "cup", wpm[0].Word)

	fmt.Println(GetOtherInBetweenPositions(line, wpm[0], wpi[0]))
}

func TestTopHat(t *testing.T) {
	vector := []float64{0, 0, 0, 1, 0, 1, 1, 0, 0, 5, 4, 2, 6, 4, 1, 0, 0, 0, 4, 0, 0}
	s, e := getBestTopHatPositions(vector)
	assert.Equal(t, 9, s)
	assert.Equal(t, 14, e)
}
