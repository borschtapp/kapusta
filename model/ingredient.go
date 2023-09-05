package model

import (
	"github.com/borschtapp/krip/utils"
)

type Ingredient struct {
	Quantity    float64 `json:"quantity,omitempty"`
	QuantityMax float64 `json:"quantityMax,omitempty"`
	Unit        string  `json:"unit,omitempty"`

	Ingredient string `json:"ingredient,omitempty"`
	Annotation string `json:"annotation,omitempty"`
}

func (r *Ingredient) String() (s string) {
	s += utils.FormatFraction(r.Quantity)
	if r.QuantityMax > 0 {
		s += "-" + utils.FormatFraction(r.QuantityMax)
	}

	if r.Unit == "" {
		s += " <unit>"
	} else {
		s += " " + r.Unit
	}

	s += " " + r.Ingredient

	if len(r.Annotation) > 0 {
		s += " (" + r.Annotation + ")"
	}
	return s
}
