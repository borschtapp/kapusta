package utils

import (
	"errors"
	"net/url"
	"regexp"
	"strconv"
	"strings"

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
