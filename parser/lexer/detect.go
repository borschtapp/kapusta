package lexer

import (
	"sync"

	"github.com/borschtapp/kapusta/dictionary"
)

// Detect runs the parse function for all supported languages and returns the one with the highest score.
// If lang is not empty, it simply returns the result for that language, including any error.
func Detect[T any](lang string, parse func(string) (T, int, error)) (T, error) {
	if lang != "" {
		res, _, err := parse(lang)
		return res, err
	}

	// Heuristic: try English first for a quick exit if it's high confidence
	enRes, enScore, enErr := parse("en")
	if enScore >= 50 && enErr == nil {
		return enRes, nil
	}

	langs := dictionary.SupportedLanguages
	type result struct {
		val   T
		score int
		err   error
	}
	results := make([]result, len(langs))
	var wg sync.WaitGroup

	for i, l := range langs {
		if l == "en" {
			results[i] = result{enRes, enScore, enErr}
			continue
		}
		wg.Add(1)
		go func(i int, l string) {
			defer wg.Done()
			val, s, err := parse(l)
			results[i] = result{val, s, err}
		}(i, l)
	}
	wg.Wait()

	best := result{enRes, enScore, enErr}

	for i, res := range results {
		l := langs[i]
		if l == "en" {
			continue
		}

		isBetter := (res.err == nil && best.err != nil) || (res.err == nil && best.err == nil && res.score > best.score)
		if isBetter {
			best = res
		}
	}

	// If no language provided any meaningful result, fallback to English
	if best.err == nil && best.score <= 0 {
		return enRes, enErr
	}

	return best.val, best.err
}
