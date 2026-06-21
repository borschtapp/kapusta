package instruction

import (
	"testing"

	"github.com/borschtapp/kapusta/model"
	"github.com/stretchr/testify/assert"
)

func TestExtractAmounts(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []model.Amount
	}{
		{
			name:  "single",
			input: "Add 1 cup flour.",
			want: []model.Amount{
				{Value: 1, Unit: "cup", Raw: "1 cup"},
			},
		},
		{
			name:  "range",
			input: "Add 1-2 cups flour.",
			want: []model.Amount{
				{Value: 1, MaxValue: 2, Unit: "cup", Raw: "1-2 cups"},
			},
		},
		{
			name:  "mixed",
			input: "Add 1 ½ cups flour.",
			want: []model.Amount{
				{Value: 1.5, Unit: "cup", Raw: "1 ½ cups"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInstruction(tt.input, Options{Lang: "en"})
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got.Amounts)
		})
	}
}

func TestExtractTemperatures(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []model.Temperature
	}{
		{
			name:  "celsius",
			input: "Preheat oven to 180°C.",
			want: []model.Temperature{
				{Value: 180, Unit: "C", Raw: "180°C"},
			},
		},
		{
			name:  "celsius no degree",
			input: "Heat to 100 C.",
			want: []model.Temperature{
				{Value: 100, Unit: "C", Raw: "100 C."},
			},
		},
		{
			name:  "fahrenheit",
			input: "Bake at 350°F.",
			want: []model.Temperature{
				{Value: 350, Unit: "F", Raw: "350°F"},
			},
		},
		{
			name:  "celsius no space",
			input: "180°C",
			want: []model.Temperature{
				{Value: 180, Unit: "C", Raw: "180°C"},
			},
		},
		{
			name:  "range",
			input: "Heat between 70-80°C.",
			want: []model.Temperature{
				{Value: 70, MaxValue: 80, Unit: "C", Raw: "70-80°C"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInstruction(tt.input, Options{Lang: "en"})
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got.Temperatures)
		})
	}
}
