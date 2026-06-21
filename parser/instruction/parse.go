package instruction

import (
	"fmt"
	"sort"
	"strings"

	"github.com/borschtapp/kapusta/dictionary"
	"github.com/borschtapp/kapusta/model"
	"github.com/borschtapp/kapusta/parser/lexer"
)

const (
	schemeTimer      = "timer"
	schemeTemp       = "temp"
	schemeAmount     = "amount"
	schemeIngredient = "ingredient"
	defaultIDPrefix  = "kap"
)

type replacement struct {
	startIdx int
	endIdx   int
	markdown string
}

type extractedQuantity struct {
	min      float64
	max      float64
	unit     lexer.Token
	startIdx int
	endIdx   int
}

func hasOverlap(start, end int, existing []replacement) bool {
	for _, r := range existing {
		if start < r.endIdx && end > r.startIdx {
			return true
		}
	}
	return false
}

func scanQuantities(tokens []lexer.Token, unitType lexer.TokenType) []extractedQuantity {
	var res []extractedQuantity
	c := &lexer.TokenCursor{Tokens: tokens}

	for !c.IsEOF() {
		startPos := c.Pos
		tok := c.Peek()
		if !tok.Type.IsNumber() {
			c.Next()
			continue
		}

		minV, maxV := parseQuantity(c)
		next := c.Peek()
		matched := false
		var overrideCode string

		if next.Type == unitType {
			matched = true
		} else if unitType == lexer.ItemTemperatureUnit && next.Type == lexer.ItemUnit {
			// Handle ambiguity: "C" might be lexed as ItemUnit (cup) but we want ItemTemperatureUnit
			l := strings.TrimSuffix(strings.ToLower(next.Lexeme), ".")
			// Heuristic: if 0 <= value > 10, it's more likely to be Celsius than cups.
			if (l == "c" && (minV <= 0 || minV > 10)) || l == "f" || l == "k" {
				matched = true
				overrideCode = strings.ToUpper(l)
			}
		}

		if matched {
			unitTok := c.Next()
			if overrideCode != "" {
				unitTok.Code = overrideCode
				unitTok.Type = lexer.ItemTemperatureUnit
			}
			res = append(res, extractedQuantity{
				min:      minV,
				max:      maxV,
				unit:     unitTok,
				startIdx: tok.StartIndex,
				endIdx:   unitTok.EndIndex,
			})
		} else {
			// Not a match for this unit type, move to the next token from where we started
			c.Pos = startPos + 1
		}
	}
	return res
}

// Options holds optional parameters for instruction parsing.
type Options struct {
	Lang               string
	Ingredients        []model.IngredientRef // input hints for name recognition
	DisableAmounts     bool                  // disable numeric+unit span annotation
	DisableTemps       bool                  // disable temperature span annotation
	DisableTimers      bool                  // disable timer span annotation
	DisableIngredients bool                  // disable ingredient matching and inference
	IDPrefix           string                // prefix for auto-generated ingredient IDs (e.g. "kap"); defaults to defaultIDPrefix
}

func ParseInstruction(text string, opts Options) (model.Instruction, error) {
	// Build the language-independent ingredient lookups once, not per detected language.
	var matcher *dictionary.Matcher
	if len(opts.Ingredients) > 0 {
		knownNames := make([]string, len(opts.Ingredients))
		for i, ref := range opts.Ingredients {
			knownNames[i] = ref.Name
		}
		matcher = dictionary.NewMatcher(knownNames)
	}

	refMap := make(map[string]model.IngredientRef)
	for _, ref := range opts.Ingredients {
		refMap[strings.ToLower(ref.Name)] = ref
	}

	return lexer.Detect(opts.Lang, func(l string) (model.Instruction, int, error) {
		o := opts
		o.Lang = l
		return parseInstruction(text, o, matcher, refMap)
	})
}

func parseInstruction(text string, opts Options, matcher *dictionary.Matcher, refMap map[string]model.IngredientRef) (model.Instruction, int, error) {
	inst := model.Instruction{
		Text:     text,
		Markdown: text,
	}

	l, err := lexer.LexWithMatcher(text, opts.Lang, matcher)
	if err != nil {
		return inst, 0, err
	}
	tokens, err := collectTokens(l)
	if err != nil {
		return inst, 0, err
	}

	var reps []replacement
	idCounter := 0

	if !opts.DisableTimers {
		timers, timerReps := extractTimers(tokens, text)
		inst.Timers = timers
		reps = append(reps, timerReps...)
	}

	if !opts.DisableTemps {
		temps, tempReps := extractTemperatures(tokens, text, reps)
		inst.Temperatures = temps
		reps = append(reps, tempReps...)
	}

	if !opts.DisableAmounts {
		amounts, amtReps := extractAmounts(tokens, text, reps)
		inst.Amounts = amounts
		reps = append(reps, amtReps...)
	}

	prefix := opts.IDPrefix
	if prefix == "" {
		prefix = defaultIDPrefix
	}

	if !opts.DisableIngredients {
		ingreds, ingReps := collectIngredients(tokens, text, refMap, &idCounter, prefix, reps)
		inst.Ingredients = ingreds
		reps = append(reps, ingReps...)

		// Fallback to infer if no ingredients matched and no hints were provided
		if len(inst.Ingredients) == 0 && len(opts.Ingredients) == 0 {
			inferred, inferReps := inferIngredients(tokens, text, &idCounter, prefix, reps)
			inst.Ingredients = inferred
			reps = append(reps, inferReps...)
		}
	}

	inst.Markdown = applyReplacements(text, reps)
	return inst, l.Score(), nil
}

// parseQuantity extracts a numeric value or range from tokens starting at i.
func parseQuantity(c *lexer.TokenCursor) (min, max float64) {
	tok := c.Peek()
	if !tok.Type.IsNumber() {
		return 0, 0
	}

	min = tok.Value
	c.Next()

	// Check for fraction after number (e.g., "1 ½")
	if c.Peek().Type == lexer.ItemNumberFraction {
		min += c.Next().Value
	}

	max = min

	// Check for range (e.g., "5-10" or "1 to 2")
	if c.Peek().Type == lexer.ItemIdentifierRange && c.PeekAt(1).Type.IsNumber() {
		c.Next() // skip range separator
		max = c.Next().Value

		// Check for fraction in the max value (e.g., "1 to 2 ½")
		if c.Peek().Type == lexer.ItemNumberFraction {
			max += c.Next().Value
		}
	}

	return min, max
}

func collectTokens(l *lexer.Lexer) ([]lexer.Token, error) {
	var tokens []lexer.Token
	for {
		tok, err := l.Next()
		if tok.Type == lexer.ItemEOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if tok.Type == lexer.ItemComment && len(tokens) > 0 {
			prev := &tokens[len(tokens)-1]
			if prev.ShouldMerge(tok) {
				prev.Lexeme += tok.Lexeme
				prev.EndIndex = tok.EndIndex
				continue
			}
		}

		tokens = append(tokens, tok)
	}
	return tokens, nil
}

func markdownLink(label, scheme, target string) string {
	if target == "" {
		return label
	}
	return fmt.Sprintf("[%s](%s:%s)", label, scheme, target)
}

// applyReplacements stitches the original text with link substitutions.
func applyReplacements(text string, reps []replacement) string {
	if len(reps) == 0 {
		return text
	}
	sort.Slice(reps, func(i, j int) bool {
		return reps[i].startIdx < reps[j].startIdx
	})

	var b strings.Builder
	b.Grow(len(text) + len(reps)*32)

	cursor := 0
	for _, r := range reps {
		if r.startIdx < cursor {
			continue // overlaps
		}
		b.WriteString(text[cursor:r.startIdx])
		b.WriteString(r.markdown)
		cursor = r.endIdx
	}
	b.WriteString(text[cursor:])
	return b.String()
}
