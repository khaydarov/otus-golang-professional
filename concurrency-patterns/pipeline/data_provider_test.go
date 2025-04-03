package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDeterministicIntStream(t *testing.T) {
	terminate := make(chan struct{})
	defer close(terminate)

	input := []int{1, 2, 3, 4, 5}
	stream := NewDeterministicIntStream(terminate, input)

	var result []int
	for v := range stream {
		result = append(result, v.(int))
	}

	assert.Equal(t, input, result)
}

func TestRandomIntStream(t *testing.T) {
	terminate := make(chan struct{})
	stream := NewRandomIntStream(terminate)

	// Test that we can get at least 3 values
	count := 0
	done := make(chan struct{})

	go func() {
		for range stream {
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
		assert.GreaterOrEqual(t, count, 3)
	case <-time.After(time.Second):
		t.Fatal("random stream did not produce enough values in time")
	}
}

func TestDeterministicInts(t *testing.T) {
	input := []int{1, 2, 3}
	fn := DeterministicInts(input)

	// Test normal values
	for _, expected := range input {
		val, err := fn()
		assert.NoError(t, err)
		assert.Equal(t, expected, val)
	}

	// Test end of stream
	val, err := fn()
	assert.Equal(t, ErrNoMoreInts, err)
	assert.Equal(t, 0, val)
}
