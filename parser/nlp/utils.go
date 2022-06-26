package nlp

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

func getWordPositions(s string, corpus []string) (wordPositions []WordPosition) {
	wordPositions = []WordPosition{}
	for _, ing := range corpus {
		pos := strings.Index(s, ing)
		if pos > -1 {
			s = strings.Replace(s, ing, strings.Repeat(" ", utf8.RuneCountInString(ing)), 1)
			ing = strings.TrimSpace(ing)
			wordPositions = append(wordPositions, WordPosition{ing, pos})
			// fmt.Println(s)
		}
	}
	sort.Slice(wordPositions, func(i, j int) bool {
		return wordPositions[i].Position < wordPositions[j].Position
	})
	return
}

// GetOtherInBetweenPositions returns the word positions comment string in the ingredients
func GetOtherInBetweenPositions(s string, pos1, pos2 WordPosition) (other string) {
	if pos1.Position > pos2.Position {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			log.Print(s, pos1, pos2)
			log.Print(r)
		}
	}()
	other = s[pos1.Position+len(pos1.Word)+1 : pos2.Position]
	other = strings.TrimSpace(other)
	return
}

// GetIngredientsInString returns the word positions of the ingredients
func GetIngredientsInString(s string) (wordPositions []WordPosition) {
	return getWordPositions(s, CorpusIngredients)
}

// GetNumbersInString returns the word positions of the numbers in the ingredients string
func GetNumbersInString(s string) (wordPositions []WordPosition) {
	return getWordPositions(s, CorpusNumbers)
}

// GetMeasuresInString returns the word positions of the measures in a ingredients string
func GetMeasuresInString(s string) (wordPositions []WordPosition) {
	return getWordPositions(s, CorpusMeasures)
}

// SanitizeLine removes parentheses, trims the line, converts to lower case,
// replaces fractions with unicode and then does special conversion for ingredients (like eggs).
func SanitizeLine(s string) string {
	s = strings.ToLower(s)
	s = strings.Replace(s, "⁄", "/", -1)
	s = strings.Replace(s, " / ", "/", -1)

	// remove parentheses
	re := regexp.MustCompile(`(?s)\((.*)\)`)
	for _, m := range re.FindAllStringSubmatch(s, -1) {
		s = strings.Replace(s, m[0], " ", 1)
	}

	s = " " + strings.TrimSpace(s) + " "

	// replace unicode fractions with fractions
	for v := range CorpusFractionNumberMap {
		s = strings.Replace(s, v, CorpusFractionNumberMap[v].FractionString, -1)
	}

	// remove non-alphanumeric
	reg, _ := regexp.Compile("[^a-zA-Z0-9/]+")
	s = reg.ReplaceAllString(s, " ")

	// replace fractions with unicode fractions
	for v := range CorpusFractionNumberMap {
		s = strings.Replace(s, CorpusFractionNumberMap[v].FractionString, v, -1)
	}

	s = strings.Replace(s, " one ", " 1 ", -1)
	s = strings.Replace(s, " egg ", " eggs ", -1)
	s = strings.Replace(s, " apples ", " apple ", -1)

	return strings.TrimSpace(s)
}

func generateHat(length, start, stop int, value float64) []float64 {
	f := make([]float64, length)
	for i := start; i < stop; i++ {
		f[i] = value
	}
	return f
}

func calculateResidual(fs1, fs2 []float64) float64 {
	res := 0.0
	if len(fs1) != len(fs2) {
		return -1
	}
	for i := range fs1 {
		res += math.Pow(fs1[i]-fs2[i], 2)
	}
	return res
}

func getBestTopHatPositions(vectorFloat []float64) (start, end int) {
	bestTopHatResidual := 1e9
	for i, v := range vectorFloat {
		if v < 2 {
			continue
		}
		for j, w := range vectorFloat {
			if j <= i || w < 1 {
				continue
			}
			hat := generateHat(len(vectorFloat), i, j, AverageFloats(vectorFloat[i:j]))
			res := calculateResidual(vectorFloat, hat) / float64(len(vectorFloat))
			if res < bestTopHatResidual {
				bestTopHatResidual = res
				start = i
				end = j
			}
		}
	}
	return
}

func convertStringToNumber(s string) float64 {
	switch s {
	case "½":
		return 0.5
	case "¼":
		return 0.25
	case "¾":
		return 0.75
	case "⅛":
		return 1.0 / 8
	case "⅜":
		return 3.0 / 8
	case "⅝":
		return 5.0 / 8
	case "⅞":
		return 7.0 / 8
	case "⅔":
		return 2.0 / 3
	case "⅓":
		return 1.0 / 3
	}
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func amountToString(amount float64) string {
	r, _ := parseDecimal(fmt.Sprintf("%2.10f", amount))
	rationalFraction := float64(r.n) / float64(r.d)
	if rationalFraction > 0 {
		bestFractionDiff := 1e9
		bestFraction := 0.0
		var fractions = map[float64]string{
			0:       "",
			1:       "",
			1.0 / 2: "1/2",
			1.0 / 3: "1/3",
			2.0 / 3: "2/3",
			1.0 / 6: "1/6",
			1.0 / 8: "1/8",
			3.0 / 8: "3/8",
			5.0 / 8: "5/8",
			7.0 / 8: "7/8",
			1.0 / 4: "1/4",
			3.0 / 4: "3/4",
		}
		for f := range fractions {
			currentDiff := math.Abs(f - rationalFraction)
			if currentDiff < bestFractionDiff {
				bestFraction = f
				bestFractionDiff = currentDiff
			}
		}
		if fractions[bestFraction] == "" {
			return strconv.FormatInt(int64(math.Round(amount)), 10)
		}
		if r.i > 0 {
			return strconv.FormatInt(r.i, 10) + " " + fractions[bestFraction]
		} else {
			return fractions[bestFraction]
		}
	}
	return strconv.FormatInt(r.i, 10)
}

func AverageFloats(fs []float64) float64 {
	f := 0.0
	for _, v := range fs {
		f += v
	}
	return f / float64(len(fs))
}

// A rational number r is expressed as the fraction p/q of two integers:
// r = p/q = (d*i+n)/d.
type rational struct {
	i int64 // integer
	n int64 // fraction numerator
	d int64 // fraction denominator
}

func gcd(x, y int64) int64 {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func parseDecimal(s string) (r rational, err error) {
	sign := int64(1)
	if strings.HasPrefix(s, "-") {
		sign = -1
	}
	p := strings.IndexByte(s, '.')
	if p < 0 {
		p = len(s)
	}
	if i := s[:p]; len(i) > 0 {
		if i != "+" && i != "-" {
			r.i, err = strconv.ParseInt(i, 10, 64)
			if err != nil {
				return rational{}, err
			}
		}
	}
	if p >= len(s) {
		p = len(s) - 1
	}
	if f := s[p+1:]; len(f) > 0 {
		n, err := strconv.ParseUint(f, 10, 64)
		if err != nil {
			return rational{}, err
		}
		d := math.Pow10(len(f))
		if math.Log2(d) > 63 {
			err = fmt.Errorf(
				"ParseDecimal: parsing %q: value out of range", f,
			)
			return rational{}, err
		}
		r.n = int64(n)
		r.d = int64(d)
		if g := gcd(r.n, r.d); g != 0 {
			r.n /= g
			r.d /= g
		}
		r.n *= sign
	}
	return r, nil
}
