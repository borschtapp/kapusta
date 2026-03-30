package util

import (
	"errors"
	"strconv"
	"strings"
)

func ParseInt(str string) (int, error) {
	str = strings.TrimSpace(str)
	if i, err := strconv.Atoi(str); err == nil {
		return i, nil
	}

	return 0, errors.New("unable to parse int from string: " + str)
}

func ParseFloat(str string) (float64, error) {
	str = strings.Replace(strings.TrimSpace(str), ",", ".", 1)
	if i, err := strconv.ParseFloat(str, 64); err == nil {
		return i, nil
	}

	return 0, errors.New("unable to parse float from string: " + str)
}
