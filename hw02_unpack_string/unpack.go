package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

type Node struct {
	R    rune
	Next *Node
}

func Unpack(s string) (string, error) {
	if s == "" {
		return s, nil
	}

	head := &Node{}
	curr := head
	for _, v := range s {
		curr.Next = &Node{R: v}
		curr = curr.Next
	}

	curr = head.Next
	if unicode.IsDigit(curr.R) {
		return "", ErrInvalidString
	}

	var result strings.Builder
	for curr != nil {
		if !unicode.IsDigit(curr.R) {
			if curr.Next != nil && unicode.IsDigit(curr.Next.R) {
				count, err := strconv.Atoi(string(curr.Next.R))
				if err == nil {
					result.WriteString(strings.Repeat(string(curr.R), count))
				} else {
					result.WriteRune(curr.R)
				}
			} else {
				result.WriteRune(curr.R)
			}
		} else if curr.Next != nil && unicode.IsDigit(curr.Next.R) {
			return "", ErrInvalidString
		}
		curr = curr.Next
	}

	return result.String(), nil
}
