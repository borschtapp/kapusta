package dictionary

import (
	"embed"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"gopkg.in/yaml.v3"
)

type Dict struct {
	Units           map[string][]string `yaml:"units"`
	QuantityBetween []string            `yaml:"quantityBetween"`
	Numbers         map[string]float64  `yaml:"numbers"`

	// Internal trie for efficient unit lookup
	trie *trieNode
}

type trieNode struct {
	children map[rune]*trieNode
	isEnd    bool
	code     string
}

func (d *Dict) buildTrie() {
	d.trie = &trieNode{children: make(map[rune]*trieNode)}
	for code, variants := range d.Units {
		for _, v := range variants {
			d.insert(v, code)
		}
	}
}

func (d *Dict) insert(str, code string) {
	node := d.trie
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

func (d *Dict) FindUnit(input string) (variant string, code string, ok bool) {
	if d.trie == nil {
		return "", "", false
	}

	node := d.trie
	pos := 0
	matchLen := 0

	// Iterate through input
	for pos < len(input) {
		r, w := utf8.DecodeRuneInString(input[pos:])
		lowerR := unicode.ToLower(r)

		if node.children == nil {
			break
		}
		next := node.children[lowerR]
		if next == nil {
			break
		}
		node = next
		pos += w

		if node.isEnd {
			// Check that the match ends on a word boundary.
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

func (d *Dict) FindNumber(str string) (float64, bool) {
	for key, val := range d.Numbers {
		if strings.EqualFold(key, str) {
			return val, true
		}
	}
	return 0, false
}

func (d *Dict) FindQuantityBetween(str string) (string, bool) {
	for _, val := range d.QuantityBetween {
		if strings.EqualFold(val, str) {
			return val, true
		}
	}
	return "", false
}

//go:embed *.yml
var fs embed.FS
var dictMap = make(map[string]*Dict)

// Called at package initialization
func init() {
	files, err := fs.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		data, err := fs.ReadFile(file.Name())
		if err != nil {
			panic(err)
		}

		var dict Dict
		err = yaml.Unmarshal(data, &dict)
		if err != nil {
			panic(err)
		}

		dict.buildTrie()

		lang := strings.Split(file.Name(), ".")[0]
		dictMap[lang] = &dict
	}
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
