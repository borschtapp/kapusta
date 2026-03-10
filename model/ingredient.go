package model

import (
	"github.com/borschtapp/krip/utils"
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
	s += utils.FormatFraction(r.Amount)
	if r.MaxAmount > 0 {
		s += "-" + utils.FormatFraction(r.MaxAmount)
	}

	if r.Unit == "" {
		s += " <unit>"
	} else {
		s += " " + r.Unit
	}

	s += " " + r.Name

	if len(r.Description) > 0 {
		s += " (" + r.Description + ")"
	}
	return s
}
