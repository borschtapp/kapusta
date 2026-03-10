# Kapusta - Recipe Parser

Kapusta is an ingredient parser written in Go. It fast and reliably parses raw recipe string inputs into structured ingredient data models, making it easier to build recipe management apps, shopping list generators, or cooking tools.

## Features

- **Natural Language Parsing**: Parses free-form ingredient text (e.g., "1 1/2 cups diced tomatoes").
- **Multi-Language Support**: Supports multiple languages using embedded YAML dictionaries.
- **Fraction and Range Support**: Correctly parses fractions (e.g., "1/2") and ranges (e.g., "1-2" or "1 to 2").
- **Extracts Annotations**: Separates the main ingredient name from additional comments or preparations (e.g., "chopped", "to taste").

## Installation

```bash
go get github.com/borschtapp/kapusta
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/borschtapp/kapusta"
)

func main() {
    ingredientLine := "1 1/2 cups diced tomatoes, drained"
    
    parsed, err := kapusta.ParseIngredient(ingredientLine, "en")
    if err != nil {
        log.Fatalf("Failed to parse: %v", err)
    }

    fmt.Printf("Quantity: %v\n", parsed.Amount)
    fmt.Printf("Unit: %s\n", parsed.Unit)
    fmt.Printf("Ingredient: %s\n", parsed.Name)
    fmt.Printf("Annotation: %s\n", parsed.Description)
}
```

## Data Model

The parsed `Ingredient` object has the following structure:

```go
type Ingredient struct {
	Amount      float64 `json:"amount,omitempty"`
	MaxAmount   float64 `json:"maxAmount,omitempty"`
	Unit        string  `json:"unit,omitempty"`
	UnitCode    string  `json:"unitCode,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
}
```
