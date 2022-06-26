package nlp

import (
	"fmt"
	"math"
	"strings"
)

// WordPosition shows a word and its position
// Note: the position is memory-dependent as it will
// be the position after the last deleted word
type WordPosition struct {
	Word     string
	Position int
}

// LineInfo has all the information for the parsing of a given line
type LineInfo struct {
	LineOriginal        string
	Line                string         `json:",omitempty"`
	IngredientsInString []WordPosition `json:",omitempty"`
	AmountInString      []WordPosition `json:",omitempty"`
	MeasureInString     []WordPosition `json:",omitempty"`

	IngredientAmount float64 `json:",omitempty"`
	IngredientName   string  `json:",omitempty"`
}

func (lineInfo *LineInfo) getTotalAmount() (err error) {
	lastPosition := -1
	totalAmount := 0.0
	wps := lineInfo.AmountInString
	for i := range wps {
		wps[i].Word = strings.TrimSpace(wps[i].Word)
		if lastPosition == -1 {
			totalAmount = convertStringToNumber(wps[i].Word)
		} else if math.Abs(float64(wps[i].Position-lastPosition)) < 6 {
			totalAmount += convertStringToNumber(wps[i].Word)
		}
		lastPosition = wps[i].Position + len(wps[i].Word)
	}
	if totalAmount == 0 && strings.Contains(lineInfo.Line, "whole") {
		totalAmount = 1
	}
	if totalAmount == 0 {
		err = fmt.Errorf("no amount found")
		lineInfo.IngredientAmount = totalAmount
	}
	return
}

func (lineInfo *LineInfo) getIngredient() (err error) {
	if len(lineInfo.IngredientsInString) == 0 {
		err = fmt.Errorf("no ingredients found")
		return
	}
	lineInfo.IngredientName = lineInfo.IngredientsInString[0].Word
	return
}
