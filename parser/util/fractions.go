package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const Fractions = "¼½¾⅐⅑⅒⅓⅔⅕⅖⅗⅘⅙⅚⅛⅜⅝⅞"

var fractionsMap = map[string]float64{
	"¼": 0.25,
	"½": 0.50,
	"¾": 0.75,
	"⅐": 0.1428571428571429,
	"⅑": 0.1111111111111111,
	"⅒": 0.1,
	"⅓": 0.3333333333333333,
	"⅔": 0.6666666666666667,
	"⅕": 0.20,
	"⅖": 0.40,
	"⅗": 0.60,
	"⅘": 0.8,
	"⅙": 0.1666666666666667,
	"⅚": 0.8333333333333333,
	"⅛": 0.125,
	"⅜": 0.375,
	"⅝": 0.625,
	"⅞": 0.875,
}

func IsFraction(r rune) bool {
	return strings.ContainsRune(Fractions, r)
}

func ParseFraction(str string) (float64, error) {
	str = strings.TrimSpace(str)
	var res float64

	// 1. Strip and sum unicode fraction symbols.
	for symbol, value := range fractionsMap {
		if strings.Contains(str, symbol) {
			str = strings.Replace(str, symbol, "", 1)
			res += value
		}
	}
	str = strings.TrimSpace(str)

	// 2. Strip and parse a trailing n/m slash fraction.
	if idx := strings.LastIndex(str, "/"); idx != -1 {
		start := strings.LastIndex(str[:idx], " ") + 1
		slashToken := str[start:]
		arr := strings.SplitN(slashToken, "/", 2)
		num, err := strconv.ParseFloat(strings.TrimSpace(arr[0]), 64)
		if err != nil {
			return 0, fmt.Errorf("unable to parse fraction from %q: %w", str, err)
		}
		den, err := strconv.ParseFloat(strings.TrimSpace(arr[1]), 64)
		if err != nil {
			return 0, fmt.Errorf("unable to parse fraction from %q: %w", str, err)
		}
		if den == 0 {
			return 0, fmt.Errorf("unable to parse fraction from %q: division by zero", str)
		}
		res += num / den
		str = strings.TrimSpace(str[:start])
	}

	// 3. Parse remaining integer/float prefix.
	if str != "" {
		val, err := ParseFloat(str)
		if err != nil {
			return 0, fmt.Errorf("unable to parse fraction from %q: %w", str, err)
		}
		res += val
	}

	return res, nil
}

func FormatFraction(f float64) string {
	integer, fraction := math.Modf(f)
	if fraction == 0 {
		return strconv.FormatInt(int64(integer), 10)
	}

	for key, value := range fractionsMap {
		if math.Abs(value-fraction) < 0.001 {
			if integer == 0 {
				return key
			}

			return strconv.FormatInt(int64(integer), 10) + " " + key
		}
	}

	return fmt.Sprintf("%v", f)
}
