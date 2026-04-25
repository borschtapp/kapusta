package model

import (
	"strings"

	"github.com/borschtapp/kapusta/parser/util"
)

type Ingredient struct {
	Amount      float64 `json:"amount,omitempty"`
	MaxAmount   float64 `json:"maxAmount,omitempty"`
	Unit        string  `json:"unit,omitempty"`
	UnitCode    string  `json:"unitCode,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
}

func (r *Ingredient) String() string {
	var b strings.Builder
	b.WriteString(util.FormatFraction(r.Amount))
	if r.MaxAmount > 0 {
		b.WriteByte('-')
		b.WriteString(util.FormatFraction(r.MaxAmount))
	}
	if r.Unit != "" {
		b.WriteByte(' ')
		b.WriteString(r.Unit)
	}
	if r.Name != "" {
		b.WriteByte(' ')
		b.WriteString(r.Name)
	}
	if r.Description != "" {
		b.WriteString(" (")
		b.WriteString(r.Description)
		b.WriteByte(')')
	}
	return b.String()
}
