package instruction

import (
	"fmt"

	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/lexer"
)

// extractTemperatures scans tokens for the pattern: number [range-sep number]? [degree-symbol]? temperature-unit
func extractTemperatures(tokens []lexer.Token, text string, existing []replacement) ([]model.Temperature, []replacement) {
	var temps []model.Temperature
	var reps []replacement

	matches := scanQuantities(tokens, lexer.ItemTemperatureUnit)
	for _, m := range matches {
		if hasOverlap(m.startIdx, m.endIdx, existing) {
			continue
		}
		minV, maxV := m.min, m.max
		if minV == maxV {
			maxV = 0
		}

		orig := text[m.startIdx:m.endIdx]
		temps = append(temps, model.Temperature{
			Value:    minV,
			MaxValue: maxV,
			Unit:     m.unit.Code,
			Raw:      orig,
		})

		target := fmt.Sprintf("%g?unit=%s", minV, m.unit.Code)
		if maxV > 0 {
			target += fmt.Sprintf("&max=%g", maxV)
		}

		reps = append(reps, replacement{
			startIdx: m.startIdx,
			endIdx:   m.endIdx,
			markdown: markdownLink(orig, schemeTemp, target),
		})
	}

	return temps, reps
}
