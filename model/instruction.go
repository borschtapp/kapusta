package model

// IngredientRef is used both as an input hint (Options.Ingredients) and as
// an output in Instruction.Ingredients. As input, Name is matched against the
// instruction text; ID is the stable backend key written into the URI.
// As output, ID is populated for all ingredients (created if not provided as input).
type IngredientRef struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

// Timer represents a time duration in seconds, extracted from an instruction.
type Timer struct {
	Value    int    `json:"value"`
	MaxValue int    `json:"max,omitempty"`
	Raw      string `json:"raw,omitempty"`
}

// Temperature represents a temperature span, extracted from an instruction.
type Temperature struct {
	Value    float64 `json:"value"`
	MaxValue float64 `json:"max,omitempty"`
	// Unit is a canonical unit code (e.g. "C", "F" or "K")
	Unit string `json:"unit,omitempty"`
	Raw  string `json:"raw,omitempty"`
}

// Amount represents a quantity of ingredients, extracted from an instruction.
type Amount struct {
	Value    float64 `json:"value"`
	MaxValue float64 `json:"max,omitempty"`
	// Unit is canonical unit code (e.g. "cup", "tbsp", "kg")
	Unit string `json:"unit,omitempty"`
	Raw  string `json:"raw,omitempty"`
}

// Instruction represents a parsed recipe instruction with extracted entities.
type Instruction struct {
	Text         string          `json:"text"`
	Markdown     string          `json:"markdown"`
	Timers       []Timer         `json:"timers"`
	Temperatures []Temperature   `json:"temperatures,omitempty"`
	Ingredients  []IngredientRef `json:"ingredients"`
	Amounts      []Amount        `json:"amounts,omitempty"`
}
