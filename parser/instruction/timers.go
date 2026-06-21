package instruction

import (
	"fmt"
	"strings"

	"github.com/borschtapp/kapusta/dictionary"
	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/lexer"
)

func buildTimerTarget(s, ms int) string {
	target := fmt.Sprintf("%d", s)
	if ms > 0 {
		target += fmt.Sprintf("?max=%d", ms)
	}
	return target
}

// extractTimers scans tokens for the pattern: number [range-sep number] time-unit
// and returns the extracted timers with their markdown replacements.
func extractTimers(tokens []lexer.Token, text string) ([]model.Timer, []replacement) {
	var timers []model.Timer
	var reps []replacement

	matches := scanQuantities(tokens, lexer.ItemTimeUnit)
	for _, m := range matches {
		mult := dictionary.TimeUnitSeconds(m.unit.Code)
		if mult == 0 {
			mult = 1
		}
		s, ms := int(m.min*float64(mult)), int(m.max*float64(mult))

		if len(timers) > 0 && strings.TrimSpace(text[reps[len(reps)-1].endIdx:m.startIdx]) == "" {
			t, r := &timers[len(timers)-1], &reps[len(reps)-1]

			currMax := t.MaxValue
			if currMax == 0 {
				currMax = t.Value
			}
			newMax := ms
			if newMax == 0 {
				newMax = s
			}

			t.Value += s
			t.MaxValue = currMax + newMax
			if t.MaxValue == t.Value {
				t.MaxValue = 0
			}

			t.Raw = text[r.startIdx:m.endIdx]
			r.endIdx, r.markdown = m.endIdx, markdownLink(t.Raw, schemeTimer, buildTimerTarget(t.Value, t.MaxValue))
			continue
		}

		if s == ms {
			ms = 0
		}
		orig := text[m.startIdx:m.endIdx]
		timers = append(timers, model.Timer{Value: s, MaxValue: ms, Raw: orig})
		reps = append(reps, replacement{
			startIdx: m.startIdx,
			endIdx:   m.endIdx,
			markdown: markdownLink(orig, schemeTimer, buildTimerTarget(s, ms)),
		})
	}
	return timers, reps
}
