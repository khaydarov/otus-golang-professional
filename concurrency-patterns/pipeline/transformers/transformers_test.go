package transformers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrementor(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "positive number",
			input:    5,
			expected: 6,
		},
		{
			name:     "zero",
			input:    0,
			expected: 1,
		},
		{
			name:     "negative number",
			input:    -5,
			expected: -4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Incrementor(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMultiplier(t *testing.T) {
	tests := []struct {
		name       string
		input      int
		multiplier int
		expected   int
	}{
		{
			name:       "multiply by 2",
			input:      5,
			multiplier: 2,
			expected:   10,
		},
		{
			name:       "multiply by 0",
			input:      5,
			multiplier: 0,
			expected:   0,
		},
		{
			name:       "multiply negative",
			input:      5,
			multiplier: -2,
			expected:   -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformer := Multiplier(tt.multiplier)
			result := transformer(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
