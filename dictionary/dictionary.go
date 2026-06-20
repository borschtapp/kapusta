//go:generate go run ./cmd/gen
package dictionary

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Dict is a per-language lookup table for units, numbers, quantity-range words, and time expressions.
type Dict struct {
	Units            map[string][]string `yaml:"units"`
	QuantityBetween  []string            `yaml:"quantity_between"`
	Numbers          map[string]float64  `yaml:"numbers"`
	SizeSuffix       []string            `yaml:"size_suffix"`
	TimeUnits        map[string][]string `yaml:"time_units"`
	TemperatureUnits map[string][]string `yaml:"temperature_units"`

	unitsTrie *trieNode
	timeTrie  *trieNode
	tempTrie  *trieNode

	numbersIdx    map[string]float64
	quantityIdx   map[string]struct{}
	sizeSuffixIdx map[string]struct{}
}

type trieNode struct {
	children map[rune]*trieNode
	isEnd    bool
	code     string
}

func (n *trieNode) add(str, code string) {
	node := n
	for _, r := range str {
		r = unicode.ToLower(r)
		if node.children == nil {
			node.children = make(map[rune]*trieNode)
		}
		if node.children[r] == nil {
			node.children[r] = &trieNode{}
		}
		node = node.children[r]
	}
	node.isEnd = true
	node.code = code
}

func (n *trieNode) find(input string) (variant string, code string, ok bool) {
	node := n
	pos := 0
	matchLen := 0

	for pos < len(input) {
		r, w := utf8.DecodeRuneInString(input[pos:])
		if node.children == nil {
			break
		}
		next := node.children[unicode.ToLower(r)]
		if next == nil {
			break
		}
		node = next
		pos += w

		if node.isEnd {
			nextR, _ := utf8.DecodeRuneInString(input[pos:])
			if pos == len(input) || !isAlphaNumeric(nextR) {
				matchLen = pos
				code = node.code
				ok = true
			}
		}
	}

	if ok {
		variant = input[:matchLen]
	}
	return
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '\'' || r == '’'
}

// Matcher provides O(n) prefix matching for a word set using a trie.
type Matcher struct {
	root *trieNode
}

// NewMatcher builds a trie-backed matcher from a word list.
func NewMatcher(words []string) *Matcher {
	return &Matcher{root: newTrie(map[string][]string{"": words})}
}

// Find returns the longest word from the set that prefixes input (with word-boundary check).
func (m *Matcher) Find(input string) (string, bool) {
	v, _, ok := m.root.find(input)
	return v, ok
}

func newTrie(data map[string][]string) *trieNode {
	if len(data) == 0 {
		return nil
	}
	root := &trieNode{children: make(map[rune]*trieNode)}
	for code, variants := range data {
		for _, v := range variants {
			root.add(v, code)
		}
	}
	return root
}

func (d *Dict) buildTrie() {
	d.unitsTrie = newTrie(d.Units)
	d.timeTrie = newTrie(d.TimeUnits)
	d.tempTrie = newTrie(d.TemperatureUnits)

	d.numbersIdx = make(map[string]float64, len(d.Numbers))
	for k, v := range d.Numbers {
		d.numbersIdx[strings.ToLower(k)] = v
	}

	d.quantityIdx = buildIdx(d.QuantityBetween)
	d.sizeSuffixIdx = buildIdx(d.SizeSuffix)
}

func buildIdx(items []string) map[string]struct{} {
	idx := make(map[string]struct{}, len(items))
	for _, v := range items {
		idx[strings.ToLower(v)] = struct{}{}
	}
	return idx
}

func (d *Dict) FindUnit(input string) (string, string, bool) {
	return d.FindTrie(input, d.unitsTrie)
}

func (d *Dict) FindTimeUnit(input string) (string, string, bool) {
	return d.FindTrie(input, d.timeTrie)
}

func (d *Dict) FindTemperatureUnit(input string) (string, string, bool) {
	return d.FindTrie(input, d.tempTrie)
}

func (d *Dict) FindTrie(input string, trie *trieNode) (string, string, bool) {
	if trie == nil {
		return "", "", false
	}
	return trie.find(input)
}

func (d *Dict) FindNumber(str string) (float64, bool) {
	v, ok := d.numbersIdx[strings.ToLower(str)]
	return v, ok
}

func (d *Dict) FindQuantityBetween(str string) (string, bool) {
	if _, ok := d.quantityIdx[strings.ToLower(str)]; ok {
		return str, true
	}
	return "", false
}

func (d *Dict) FindSizeSuffix(str string) (string, bool) {
	if _, ok := d.sizeSuffixIdx[strings.ToLower(str)]; ok {
		return str, true
	}
	return "", false
}

var timeUnitSeconds = map[string]int{
	"second": 1,
	"minute": 60,
	"hour":   3600,
	"day":    86400,
}

// TimeUnitSeconds returns the number of seconds for a given time unit code.
func TimeUnitSeconds(code string) int {
	return timeUnitSeconds[code]
}

func ForLanguage(lang string) (*Dict, error) {
	dict := dictMap[lang]
	if dict == nil && len(lang) > 2 {
		dict = dictMap[lang[:2]]
	}

	if dict == nil {
		return nil, fmt.Errorf("no dictionary for language %q", lang)
	}

	return dict, nil
}

var SupportedLanguages []string

func init() {
	SupportedLanguages = make([]string, 0, len(dictMap))
	for l := range dictMap {
		SupportedLanguages = append(SupportedLanguages, l)
	}
	sort.Strings(SupportedLanguages)
}
