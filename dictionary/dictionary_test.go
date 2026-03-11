package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLongestMatch(t *testing.T) {
	dict := &Dict{
		Units: map[string][]string{
			"tablespoon": {"столова ложка", "столові ложки", "ст. л."},
			"cup":        {"склянка", "склянки", "чашка", "мірна чашка"},
		},
	}
	dict.buildTrie()

	tests := []struct {
		input   string
		variant string
		code    string
		ok      bool
	}{
		{"столова ложка морської солі", "столова ложка", "tablespoon", true},
		{"мірна чашка води", "мірна чашка", "cup", true},
		{"склянка води", "склянка", "cup", true},
		{"ст. л. солі", "ст. л.", "tablespoon", true},
		{"тарілка супу", "", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			variant, code, ok := dict.FindUnit(tt.input)
			assert.Equal(t, tt.variant, variant)
			assert.Equal(t, tt.code, code)
			assert.Equal(t, tt.ok, ok)
		})
	}
}
