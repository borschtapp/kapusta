package util

import (
	"errors"
	"strconv"
	"strings"
)

func ParseFloat(str string) (float64, error) {
	str = strings.Replace(strings.TrimSpace(str), ",", ".", 1)
	if i, err := strconv.ParseFloat(str, 64); err == nil {
		return i, nil
	}

	return 0, errors.New("unable to parse float from string: " + str)
}
