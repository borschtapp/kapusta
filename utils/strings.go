package utils

import (
	"errors"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/channelmeter/iso8601duration"
	"golang.org/x/net/publicsuffix"
)

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

func RemoveNewLines(s string) string {
	return strings.Replace(strings.Replace(s, "\n", "", -1), "\r", "", -1)
}

func GetDurationMinutes(s string) (res int, ok bool) {
	d, err := duration.FromString(s)
	if err == nil {
		return int(d.ToDuration().Minutes()), true
	}

	return
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
