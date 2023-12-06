package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

// MyStringBuilder Custom string Builder to extend its functionality.
type MyStringBuilder struct {
	strings.Builder
}

// DeleteLast Removes last byte from string.
func (b *MyStringBuilder) DeleteLast() {
	s := b.String()
	s = s[:len(s)-1]
	b.Reset()
	b.WriteString(s)
}

// RepeatLast Repeats last byte `count` times.
func (b *MyStringBuilder) RepeatLast(count int) {
	s := b.String()
	repeatedLastChar := strings.Repeat(string(s[len(s)-1]), count)
	b.WriteString(repeatedLastChar)
}

func Unpack(s string) (string, error) {
	if s == "" {
		return s, nil
	}

	if isDigit(s[0]) {
		return "", ErrInvalidString
	}

	result := MyStringBuilder{}
	for i := 0; i < len(s); i++ {
		if i > 0 && isDigit(s[i]) && isDigit(s[i-1]) {
			return "", ErrInvalidString
		}

		if isDigit(s[i]) {
			digit := convertDigit(s[i])
			if digit == 0 {
				result.DeleteLast()
			} else {
				result.RepeatLast(digit - 1)
			}
		} else {
			result.WriteByte(s[i])
		}
	}

	return result.String(), nil
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func convertDigit(b byte) int {
	return int(b - '0')
}
