package main

import (
	"testing"
	"time"

	"github.com/khaydarov/otus-golang-professional/concurrency-patterns/pipeline/transformers"
	"github.com/stretchr/testify/assert"
)

func TestPipeline_Execute(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fns      []transformers.Transformer
		expected []interface{}
	}{
		{
			name:  "increment and multiply by 3",
			input: []int{1, 2, 3},
			fns: []transformers.Transformer{
				transformers.Incrementor,
				transformers.Multiplier(3),
			},
			expected: []interface{}{6, 9, 12},
		},
		{
			name:  "only increment",
			input: []int{1, 2, 3},
			fns: []transformers.Transformer{
				transformers.Incrementor,
			},
			expected: []interface{}{2, 3, 4},
		},
		{
			name:     "no transformers",
			input:    []int{1, 2, 3},
			fns:      []transformers.Transformer{},
			expected: []interface{}{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			terminate := make(chan struct{})
			defer close(terminate)

			pipeline := NewPipeline(terminate)
			result := pipeline.Execute(
				NewDeterministicIntStream(terminate, tt.input),
				tt.fns...,
			)

			var actual []interface{}
			for v := range result {
				actual = append(actual, v)
			}

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestPipeline_Termination(t *testing.T) {
	terminate := make(chan struct{})
	pipeline := NewPipeline(terminate)

	// Create an infinite stream of numbers
	result := pipeline.Execute(
		NewRandomIntStream(terminate),
		transformers.Incrementor,
	)

	// Read a few values
	count := 0
	done := make(chan struct{})

	go func() {
		for range result {
			count++
			if count >= 3 {
				close(terminate)
				break
			}
		}
		close(done)
	}()

	select {
	case <-done:
		// Test passed
	case <-time.After(time.Second):
		t.Fatal("pipeline did not terminate in time")
	}
}
