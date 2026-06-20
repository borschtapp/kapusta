# Kapusta - Recipe Parser

Kapusta is a Go library for parsing recipe text into structured data. Today it handles two main tasks:

- **Ingredient parsing** — turns `"1 1/2 cups diced tomatoes, drained"` into a typed `Ingredient` struct.
- **Instruction parsing** — extracts timers and ingredient references from a recipe step and returns annotated Markdown ready for interactive UIs.

## Features

- **Natural language parsing**: free-form ingredient strings with fractions, ranges, word numbers, and annotations.
- **Instruction annotation**: extracts `Timer` and `Ingredient` entities from step text and embeds them as Markdown links using custom URI schemes (`timer:`, `ingredient:`, `amount:`, `temp:`).
- **Multi-language support**: embedded YAML dictionaries for English, German, Spanish, French, Italian, and Ukrainian.
- **Fraction and range support**: `1/2`, `1 ½`, `1-2`, `1 to 2`, `1–2` (en-dash), word numbers (`two cups`).

## Installation

```bash
go get github.com/borschtapp/kapusta
```

## Quick Start

### Parsing an ingredient

```go
ing, err := kapusta.ParseIngredient("1 1/2 cups diced tomatoes, drained")
if err != nil {
    log.Fatal(err)
}
// ing.Amount      → 1.5
// ing.MaxAmount   → 0 (only set for ranges like "1-2 cups")
// ing.Unit        → "cups"
// ing.UnitCode    → "cup"  (canonical code from the dictionary)
// ing.Name        → "diced tomatoes"
// ing.Description → "drained"
```

### Parsing a recipe instruction

```go
import "github.com/borschtapp/kapusta/model"

inst, err := kapusta.ParseInstruction("Bake potatoes for 30 minutes.", kapusta.InstructionOptions{
    Lang:        "en",
    Ingredients: []model.IngredientRef{{Name: "potatoes"}}, // optional; omit to infer from quantity+unit context
})
if err != nil {
    log.Fatal(err)
}
// inst.Text     → "Bake potatoes for 30 minutes."
// inst.Markdown → "Bake [potatoes](ingredient:kap0) for [30 minutes](timer:1800)."
//
// inst.Timers[0].Value  → 1800
// inst.Timers[0].Raw    → "30 minutes"
//
// inst.Ingredients[0].ID   → "kap0"
// inst.Ingredients[0].Name → "potatoes"
```

When `Ingredients` is omitted, the parser infers ingredient names from tokens that immediately follow a quantity+unit (`2 tbsp sugar` → `sugar`).

### Markdown output format

`ParseInstruction` embeds extracted entities as standard CommonMark links with custom URI schemes — any renderer shows readable text while the URI carries machine-readable data:

| Example link                       | URI scheme                | Meaning               |
|------------------------------------|---------------------------|-----------------------|
| `[30 minutes](timer:1800)`         | `timer:<seconds>`         | single timer duration |
| `[2-3 minutes](timer:120?max=180)` | `timer:<min>?max=<max>`   | timer range           |
| `[potatoes](ingredient:kap0)`      | `ingredient:<id>`         | ingredient reference  |
| `[1 cup](amount:1?unit=cup)`       | `amount:<value>?unit=<u>` | measured quantity     |
| `[100°C](temp:100?unit=C)`         | `temp:<value>?unit=<u>`   | temperature           |
