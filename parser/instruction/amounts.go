package instruction

import (
	"fmt"

	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/lexer"
)

// extractAmounts scans tokens for the pattern: number [range-sep number]? unit
func extractAmounts(tokens []lexer.Token, text string, existing []replacement) ([]model.Amount, []replacement) {
	var amounts []model.Amount
	var reps []replacement

	matches := scanQuantities(tokens, lexer.ItemUnit)
	for _, m := range matches {
		if hasOverlap(m.startIdx, m.endIdx, existing) {
			continue
		}
		minV, maxV := m.min, m.max
		if minV == maxV {
			maxV = 0
		}

		orig := text[m.startIdx:m.endIdx]
		amounts = append(amounts, model.Amount{
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
			markdown: markdownLink(orig, schemeAmount, target),
		})
	}

	return amounts, reps
}
