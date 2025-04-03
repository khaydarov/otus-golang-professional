package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStage(t *testing.T) {
	terminate := make(chan struct{})
	defer close(terminate)

	// Create a simple transformer that doubles the input
	doubler := func(v interface{}) interface{} {
		return v.(int) * 2
	}

	// Create input channel
	in := make(chan interface{})
	stage := newStage(terminate, "doubler", doubler)
	out := stage(in)

	// Test values
	go func() {
		in <- 1
		in <- 2
		in <- 3
		close(in)
	}()

	expected := []int{2, 4, 6}
	var result []int
	for v := range out {
		result = append(result, v.(int))
	}

	assert.Equal(t, expected, result)
}

func TestStage_Termination(t *testing.T) {
	terminate := make(chan struct{})

	// Create a transformer that simulates work
	sleeper := func(v interface{}) interface{} {
		time.Sleep(10 * time.Millisecond)
		return v
	}

	in := make(chan interface{})
	stage := newStage(terminate, "sleeper", sleeper)
	out := stage(in)

	// Start sending values
	go func() {
		for i := 0; i < 100; i++ {
			in <- i
		}
	}()

	// Read a few values then terminate
	count := 0
	done := make(chan struct{})

	go func() {
		for range out {
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
		t.Fatal("stage did not terminate in time")
	}
}
