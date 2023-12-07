package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if s == "" {
		return s, nil
	}

	var (
		result strings.Builder
		i      int
		prev   rune
		curr   rune
	)

	for i, curr = range s {
		if (i == 0 && unicode.IsDigit(curr)) || (unicode.IsDigit(prev) && unicode.IsDigit(curr)) {
			return "", ErrInvalidString
		}

		if i == 0 {
			prev = curr
			continue
		}

		if unicode.IsDigit(curr) {
			digit, err := strconv.Atoi(string(curr))
			if err == nil {
				r := strings.Repeat(string(prev), digit)
				result.WriteString(r)
			}
		} else {
			if !unicode.IsDigit(prev) {
				result.WriteRune(prev)
			}
		}

		prev = curr
	}

	if !unicode.IsDigit(curr) {
		result.WriteRune(curr)
	}

	return result.String(), nil
}
