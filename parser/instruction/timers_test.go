package instruction

import (
	"testing"

	"github.com/borschtapp/kapusta/parser/lexer"
	"github.com/stretchr/testify/assert"
)

func TestExtractTimers(t *testing.T) {
	tests := []struct {
		text    string
		wantMin []int
		wantMax []int
	}{
		{"bake for 30 minutes", []int{1800}, []int{0}},
		{"cook 2-3 hours", []int{7200}, []int{10800}},
		{"rest for 45 seconds", []int{45}, []int{0}},
		{"simmer 1 day", []int{86400}, []int{0}},
		{"bake for 30 minutes and let it rest for 5-10 minutes", []int{1800, 300}, []int{0, 600}},
		{"proof for 1½ hours", []int{5400}, []int{0}},
		{"proof for 1 ½ hours", []int{5400}, []int{0}},
		{"bake for 1 hour 30 minutes", []int{5400}, []int{0}},
		{"no time here", nil, nil},
		{"add 2 cups flour", nil, nil}, // unit but not a time unit
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			l, err := lexer.Lex(tt.text, "en")
			assert.NoError(t, err)
			tokens, err := collectTokens(l)
			assert.NoError(t, err)
			timers, reps := extractTimers(tokens, tt.text)
			assert.Len(t, timers, len(tt.wantMin))
			assert.Len(t, reps, len(tt.wantMin))
			for i, timer := range timers {
				assert.Equal(t, tt.wantMin[i], timer.Value)
				assert.Equal(t, tt.wantMax[i], timer.MaxValue)
			}
		})
	}
}
