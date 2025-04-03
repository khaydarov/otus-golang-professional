package main

import (
	"sort"
	"testing"
)

func TestFanIn(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	// Feed data into channels
	go func() {
		ch1 <- "a1"
		ch1 <- "a2"
		close(ch1)
	}()

	go func() {
		ch2 <- "b1"
		ch2 <- "b2"
		close(ch2)
	}()

	go func() {
		ch3 <- "c1"
		ch3 <- "c2"
		close(ch3)
	}()

	// Fan in the channels
	result := fanIn(ch1, ch2, ch3)

	// Collect results
	var received []string
	for s := range result {
		received = append(received, s)
	}

	// Sort for deterministic comparison
	sort.Strings(received)
	expected := []string{"a1", "a2", "b1", "b2", "c1", "c2"}

	if len(received) != len(expected) {
		t.Errorf("Expected %d items, got %d", len(expected), len(received))
	}

	for i, v := range expected {
		if received[i] != v {
			t.Errorf("Expected %v at position %d, got %v", v, i, received[i])
		}
	}
}

func TestFanInEmpty(t *testing.T) {
	result := fanIn()

	count := 0
	for range result {
		count++
	}

	if count != 0 {
		t.Errorf("Expected 0 items from empty fan-in, got %d", count)
	}
}

func TestFanInSingleChannel(t *testing.T) {
	ch := make(chan string)
	go func() {
		ch <- "test"
		close(ch)
	}()

	result := fanIn(ch)

	received := <-result
	if received != "test" {
		t.Errorf("Expected 'test', got '%s'", received)
	}

	_, open := <-result
	if open {
		t.Error("Channel should be closed after receiving all items")
	}
}

func TestFanInClosedChannels(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	close(ch1)
	close(ch2)

	result := fanIn(ch1, ch2)

	count := 0
	for range result {
		count++
	}

	if count != 0 {
		t.Errorf("Expected 0 items from closed channels, got %d", count)
	}
}
