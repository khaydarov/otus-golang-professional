package main

import (
	"testing"
	"time"
)

func TestFanOut(t *testing.T) {
	// Create a channel with known tasks
	tasks := make(chan Task)
	go func() {
		defer close(tasks)
		tasks <- Task{ID: 1, Duration: 100 * time.Millisecond}
		tasks <- Task{ID: 2, Duration: 50 * time.Millisecond}
		tasks <- Task{ID: 3, Duration: 75 * time.Millisecond}
	}()

	// Test with 2 workers
	numWorkers := 2
	results := fanOut(numWorkers, tasks)

	// Check if correct number of result channels created
	if len(results) != numWorkers {
		t.Errorf("Expected %d result channels, got %d", numWorkers, len(results))
	}

	// Wait for all workers to complete
	for _, result := range results {
		<-result
	}
}

func TestFanOutNoTasks(t *testing.T) {
	// Test with empty task channel
	tasks := make(chan Task)
	close(tasks)

	numWorkers := 3
	results := fanOut(numWorkers, tasks)

	if len(results) != numWorkers {
		t.Errorf("Expected %d result channels, got %d", numWorkers, len(results))
	}

	// All workers should finish immediately since there are no tasks
	for _, result := range results {
		<-result
	}
}

func TestFanOutSingleWorker(t *testing.T) {
	tasks := make(chan Task)
	go func() {
		defer close(tasks)
		tasks <- Task{ID: 1, Duration: 10 * time.Millisecond}
	}()

	results := fanOut(1, tasks)

	if len(results) != 1 {
		t.Errorf("Expected 1 result channel, got %d", len(results))
	}

	<-results[0]
}
