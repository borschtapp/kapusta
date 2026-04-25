package model

import (
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

func (r *Ingredient) String() (s string) {
	s += util.FormatFraction(r.Amount)
	if r.MaxAmount > 0 {
		s += "-" + util.FormatFraction(r.MaxAmount)
	}

	if r.Unit != "" {
		s += " " + r.Unit
	}

	s += " " + r.Name

	if r.Description != "" {
		s += " (" + r.Description + ")"
	}
	return s
}
