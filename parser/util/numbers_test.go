package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFloat(t *testing.T) {
	tests := []struct {
		give    string
		want    float64
		wantErr bool
	}{
		{"1", 1, false},
		{"1.256", 1.256, false},
		{"1,35", 1.35, false},
		{"0", 0, false},
		{"-1.5", -1.5, false},
		{"  1.5  ", 1.5, false},
		{"test", 0, true},
		{"", 0, true},
		{"1,2,3", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := ParseFloat(tt.give)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
