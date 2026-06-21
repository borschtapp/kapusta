package instruction

import (
	"fmt"

	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/lexer"
)

func extractAmounts(tokens []lexer.Token, text string, existing []replacement) ([]model.Amount, []replacement) {
	return extractSpans(tokens, text, lexer.ItemUnit, schemeAmount, existing,
		func(min, max float64, code, raw string) model.Amount {
			return model.Amount{Value: min, MaxValue: max, Unit: code, Raw: raw}
		})
}

func extractTemperatures(tokens []lexer.Token, text string, existing []replacement) ([]model.Temperature, []replacement) {
	return extractSpans(tokens, text, lexer.ItemTemperatureUnit, schemeTemp, existing,
		func(min, max float64, code, raw string) model.Temperature {
			return model.Temperature{Value: min, MaxValue: max, Unit: code, Raw: raw}
		})
}

// buildQuantityTarget formats the URI target for amount/temperature links.
func buildQuantityTarget(value, max float64, unitCode string) string {
	target := fmt.Sprintf("%g?unit=%s", value, unitCode)
	if max > 0 {
		target += fmt.Sprintf("&max=%g", max)
	}
	return target
}

// extractSpans is the shared core for extractAmounts and extractTemperatures.
func extractSpans[T any](tokens []lexer.Token, text string, unitType lexer.TokenType, scheme string, existing []replacement, build func(min, max float64, code, raw string) T) ([]T, []replacement) {
	var results []T
	var reps []replacement

	for _, m := range scanQuantities(tokens, unitType) {
		if hasOverlap(m.startIdx, m.endIdx, existing) {
			continue
		}
		minV, maxV := m.min, m.max
		if minV == maxV {
			maxV = 0
		}

		orig := text[m.startIdx:m.endIdx]
		results = append(results, build(minV, maxV, m.unit.Code, orig))
		reps = append(reps, replacement{
			startIdx: m.startIdx,
			endIdx:   m.endIdx,
			markdown: markdownLink(orig, scheme, buildQuantityTarget(minV, maxV, m.unit.Code)),
		})
	}

	return results, reps
}
