package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/channelmeter/iso8601duration"
)

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
