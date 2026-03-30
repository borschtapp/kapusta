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

func ParseFraction(str string) (float64, error) {
	var res float64 = 0

	str = strings.TrimSpace(str)
	if strings.Contains(str, "/") {
		original := str
		intSplit := strings.Split(str, " ")
		str = ""
		frac := intSplit[0]
		if len(intSplit) == 2 {
			str = intSplit[0]
			frac = intSplit[1]
		} else if len(intSplit) > 2 {
			return 0, fmt.Errorf("unable to parse fractions from string `%s`: too many spaces", original)
		}

		arr := strings.Split(frac, "/")
		if len(arr) == 2 {
			if num, err := strconv.ParseFloat(arr[0], 64); err == nil {
				if den, err := strconv.ParseFloat(arr[1], 64); err == nil {
					if den == 0 {
						return 0, fmt.Errorf("unable to parse fractions from string `%s`: division by zero", original)
					}
					res += num / den
				} else {
					return 0, fmt.Errorf("unable to parse fractions from string `%s`: %w", str, err)
				}
			} else {
				return 0, fmt.Errorf("unable to parse fractions from string `%s`: %w", str, err)
			}
		} else {
			return 0, fmt.Errorf("unable to parse fractions from string `%s`: too many slashes", str)
		}
	} else if strings.ContainsAny(str, Fractions) {
		for symbol, value := range fractionsMap {
			if strings.Contains(str, symbol) {
				str = strings.Replace(str, symbol, "", 1)
				res += value
			}
		}
	}

	if len(str) > 0 {
		if val, err := ParseFloat(str); err == nil {
			res += val
		} else {
			return 0, fmt.Errorf("unable to parse fractions from string `%s`: %w", str, err)
		}
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
