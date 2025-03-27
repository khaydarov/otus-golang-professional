package hw02unpackstring

import (
	"errors"

	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UnpackSuite struct {
	suite.Suite
}

func (s *UnpackSuite) TestUnpack() {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "ac4.", expected: "acccc."},
		{input: "bananaa0", expected: "banana"},
		{input: "co2l", expected: "cool"},
		{input: "go2g1le", expected: "google"},
		{input: "cб2l", expected: "cббl"},
		{input: "a2bccdб", expected: "aabccdб"},
		{input: "a2bccdб2", expected: "aabccdбб"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc

		s.Run(tc.input, func() {
			result, err := Unpack(tc.input)
			require.NoError(s.T(), err)
			require.Equal(s.T(), tc.expected, result)
		})
	}
}

func (s *UnpackSuite) TestUnpackInvalidString() {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		s.Run(tc, func() {
			_, err := Unpack(tc)
			require.Truef(s.T(), errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestUnpackSuite(t *testing.T) {
	suite.Run(t, new(UnpackSuite))
}
