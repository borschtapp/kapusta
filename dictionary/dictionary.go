//go:generate go run ./cmd/gen
package dictionary

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Dict is a per-language lookup table for units, numbers, quantity-range words, and time expressions.
type Dict struct {
	Units           map[string][]string `yaml:"units"`
	QuantityBetween []string            `yaml:"quantity_between"`
	Numbers         map[string]float64  `yaml:"numbers"`
	SizeSuffix      []string            `yaml:"size_suffix"`

	unitsTrie *trieNode

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
			r, _ := utf8.DecodeRuneInString(input[pos:])
			if pos == len(input) || (!unicode.IsLetter(r) && !unicode.IsDigit(r)) {
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

// Matcher provides O(n) prefix matching for a word set using a trie.
type Matcher struct {
	root *trieNode
}

// NewMatcher builds a trie-backed matcher from a word list.
func NewMatcher(words []string) *Matcher {
	m := &Matcher{root: &trieNode{children: make(map[rune]*trieNode)}}
	for _, w := range words {
		m.root.add(w, "")
	}
	return m
}

// Find returns the longest word from the set that prefixes input (with word-boundary check).
func (m *Matcher) Find(input string) (string, bool) {
	v, _, ok := m.root.find(input)
	return v, ok
}

func (d *Dict) buildTrie() {
	d.unitsTrie = &trieNode{children: make(map[rune]*trieNode)}
	for code, variants := range d.Units {
		for _, v := range variants {
			d.unitsTrie.add(v, code)
		}
	}

	d.numbersIdx = make(map[string]float64, len(d.Numbers))
	for k, v := range d.Numbers {
		d.numbersIdx[strings.ToLower(k)] = v
	}

	d.quantityIdx = make(map[string]struct{}, len(d.QuantityBetween))
	for _, v := range d.QuantityBetween {
		d.quantityIdx[strings.ToLower(v)] = struct{}{}
	}

	d.sizeSuffixIdx = make(map[string]struct{}, len(d.SizeSuffix))
	for _, v := range d.SizeSuffix {
		d.sizeSuffixIdx[strings.ToLower(v)] = struct{}{}
	}
}

func (d *Dict) FindUnit(input string) (string, string, bool) {
	return d.unitsTrie.find(input)
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

func ForLang(lang string) (*Dict, error) {
	dict := dictMap[lang]
	if dict == nil && len(lang) > 2 {
		dict = dictMap[lang[:2]]
	}

	if dict == nil {
		return nil, fmt.Errorf("no dictionary for language %q", lang)
	}

	return dict, nil
}
