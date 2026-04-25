package util

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseFloat(str string) (float64, error) {
	str = strings.TrimSpace(str)
	str = strings.Replace(str, ",", ".", 1)
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to parse float from %q: %w", str, err)
	}
	return val, nil
}
