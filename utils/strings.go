package utils

import (
	"errors"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/publicsuffix"
)

var paragraphsRegex = regexp.MustCompile(`\n\r?\s*\n\r?`)

func Hostname(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		return ""
	}

	host := strings.ToLower(u.Hostname())
	host = strings.TrimPrefix(host, "www.")
	return host
}

func ParserAlias(urlStr string) string {
	alias := Hostname(urlStr)

	// remove public domain
	suffix, _ := publicsuffix.PublicSuffix(alias)
	alias = strings.TrimSuffix(alias, "."+suffix)

	// remove common prefixes
	alias = strings.TrimPrefix(alias, "api.")

	// replace dots with underscores
	alias = strings.ReplaceAll(alias, ".", "_")
	return alias
}

func Cleanup(s string) string {
	s = html.UnescapeString(s)
	s = strings.TrimSpace(s)
	return s
}

func CleanupInline(s string) string {
	s = html.UnescapeString(s)
	s = strings.ReplaceAll(s, " ", " ")
	s = strings.Join(strings.Fields(s), " ") // remove redundant spaces, but also all '\t', '\n', '\r'
	s = strings.ReplaceAll(s, " ,", ",")
	s = strings.Trim(s, "\"")
	// TODO: Remove &, written as \u0026 (bettycrocker, in images as well claudia.abril)
	// TODO: Remove >, written as \u003e (Cookstr, countryliving)
	return s
}

func RemoveSpaces(s string) string {
	return strings.Join(strings.Fields(s), "")
}

func SplitParagraphs(s string) []string {
	// TODO: check for colon, add it as a section (bigoven, blueapron)
	split := paragraphsRegex.Split(s, -1)

	var result []string
	for _, p := range split {
		p = CleanupInline(p)
		if len(p) != 0 {
			result = append(result, p)
		}
	}

	return result
}

func RemoveNewLines(s string) string {
	return strings.Replace(strings.Replace(s, "\n", "", -1), "\r", "", -1)
}

var timeRegex = regexp.MustCompile(`(?i)(\D*(?P<hours>[\d.\s/?¼½¾⅓⅔⅕⅖⅗]+)\s*(hours|hrs|hr|h|óra))?(\D*(?P<minutes>[\d.]+)\s*(minutes|mins|min|m|perc))?`)

func ParseDuration(str string) (time.Duration, bool) {
	matches := timeRegex.FindStringSubmatch(str)
	if len(matches) == 0 {
		log.Println("unable to parse duration from string: " + str)
		return 0, false
	}

	var duration time.Duration
	if hours, err := ParseFractions(matches[2]); err == nil && hours > 0 {
		duration += time.Duration(hours) * time.Hour
	}
	if minutes, err := strconv.ParseFloat(matches[5], 32); err == nil && minutes > 0 {
		duration += time.Duration(minutes) * time.Minute
	}
	return duration, true
}

func ParseInt(str string) (int, error) {
	re := regexp.MustCompile("\\d+")
	arr := re.FindAllString(str, 1)
	for _, element := range arr {
		if i, err := strconv.Atoi(element); err == nil {
			return i, nil
		}
	}

	return 0, errors.New("unable to parse int from string: " + str)
}

var fractions = map[string]float32{
	"⅛": 0.125,
	"⅕": 0.20,
	"¼": 0.25,
	"⅓": 0.33,
	"⅜": 0.375,
	"⅖": 0.40,
	"½": 0.50,
	"⅗": 0.60,
	"⅝": 0.625,
	"⅔": 0.66,
	"¾": 0.75,
	"⅘": 0.8,
	"⅞": 0.875,
}

func ParseFractions(str string) (float32, error) {
	var res float32 = 0

	if strings.Contains(str, "/") {
		intSplit := strings.Split(str, " ")
		frac := intSplit[0]
		if len(intSplit) == 2 {
			str = intSplit[0]
			frac = intSplit[1]
		} else if len(intSplit) > 2 {
			return 0, errors.New("unable to parse fractions from string `" + str + "`: too many spaces")
		}

		arr := strings.Split(frac, "/")
		if len(arr) == 2 {
			if num, err := strconv.ParseFloat(arr[0], 32); err == nil {
				if den, err := strconv.ParseFloat(arr[1], 32); err == nil {
					res += float32(num / den)
				} else {
					return 0, errors.New("unable to parse fractions from string `" + str + "`: " + err.Error())
				}
			} else {
				return 0, errors.New("unable to parse fractions from string `" + str + "`: " + err.Error())
			}
		} else {
			return 0, errors.New("unable to parse fractions from string `" + str + "`: too many slashes")
		}
	} else {
		for symbol, value := range fractions {
			if strings.Contains(str, symbol) {
				str = strings.Replace(str, symbol, "", 1)
				res += value
			}
		}
	}

	if val, err := strconv.ParseFloat(str, 32); err == nil {
		res += float32(val)
	} else {
		return 0, errors.New("unable to parse fractions from string `" + str + "`: " + err.Error())
	}

	return res, nil
}

func FindNumber(str string) int {
	split := strings.Fields(str)
	for _, s := range split {
		if i, err := strconv.Atoi(s); err == nil {
			return i
		}
	}
	return 0
}
