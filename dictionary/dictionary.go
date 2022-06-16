package dictionary

import (
	"embed"
	"errors"
	"strings"

	"gopkg.in/yaml.v3"
)

type Dict struct {
	Units           map[string][]string `yaml:"units"`
	QuantityBetween []string            `yaml:"quantityBetween"`
	Numbers         map[string]float64  `yaml:"numbers"`
}

func (d Dict) FindUnit(str string) (string, bool) {
	for _, variants := range d.Units {
		for _, variant := range variants {
			if strings.EqualFold(variant, str) {
				return variant, true
			}
		}
	}
	return "", false
}

func (d Dict) FindNumber(str string) (string, bool) {
	for key := range d.Numbers {
		if strings.EqualFold(key, str) {
			return key, true
		}
	}
	return "", false
}

func (d Dict) FindQuantityBetween(str string) (string, bool) {
	for _, val := range d.QuantityBetween {
		if strings.EqualFold(val, str) {
			return val, true
		}
	}
	return "", false
}

//go:embed *.yml
var fs embed.FS
var languages []string
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

		lang := strings.Split(file.Name(), ".")[0]
		dictMap[lang] = &dict
		languages = append(languages, lang)
	}
}

func ForLang(lang string) (*Dict, error) {
	dict := dictMap[lang]
	if dict == nil && len(lang) > 2 {
		dict = dictMap[lang[:2]]
	}

	if dict == nil {
		return nil, errors.New("no dictionary for language `" + lang + "`")
	}

	return dict, nil
}

func HasDictionary(lang string) bool {
	if lang == "" {
		return false
	}

	for _, l := range languages {
		if l == lang || l == lang[:2] {
			return true
		}
	}

	return false
}

func Default() *Dict {
	return dictMap["en"]
}
