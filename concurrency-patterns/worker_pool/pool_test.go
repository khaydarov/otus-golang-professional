package main

import (
	"testing"
	"time"
)

func TestNewWorkerPool(t *testing.T) {
	tests := []struct {
		name      string
		count     int
		wantCount int
	}{
		{"single worker", 1, 1},
		{"multiple workers", 5, 5},
		{"zero workers", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pool := NewWorkerPool(tt.count)
			if len(pool.workers) != tt.wantCount {
				t.Errorf("NewWorkerPool(%d) got %d workers, want %d", tt.count, len(pool.workers), tt.wantCount)
			}
		})
	}
}

func TestWorkerPool_Handle(t *testing.T) {
	tests := []struct {
		name      string
		workers   int
		tasks     []Task
		wantCount int
	}{
		{
			name:    "single task single worker",
			workers: 1,
			tasks: []Task{
				{id: 1, duration: 10 * time.Millisecond},
			},
			wantCount: 1,
		},
		{
			name:    "multiple tasks multiple workers",
			workers: 3,
			tasks: []Task{
				{id: 1, duration: 10 * time.Millisecond},
				{id: 2, duration: 20 * time.Millisecond},
				{id: 3, duration: 30 * time.Millisecond},
			},
			wantCount: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pool := NewWorkerPool(tt.workers)
			tasks := make(chan Task, len(tt.tasks))
			stop := make(chan struct{})
			results := pool.Handle(stop, tasks)

			// Send tasks
			for _, task := range tt.tasks {
				tasks <- task
			}
			close(tasks)

			// Create a done channel to signal when all results are collected
			done := make(chan struct{})
			go func() {
				defer close(done)
				count := 0
				resultMap := make(map[int]struct{})

				for result := range results {
					if result.err != nil {
						t.Errorf("unexpected error: %v", result.err)
					}
					resultMap[result.taskId] = struct{}{}
					count++
				}

				if count != tt.wantCount {
					t.Errorf("got %d results, want %d", count, tt.wantCount)
				}

				// Check all tasks were processed
				for _, task := range tt.tasks {
					if _, ok := resultMap[task.id]; !ok {
						t.Errorf("task %d was not processed", task.id)
					}
				}
			}()

			// Wait for results with timeout
			select {
			case <-done:
				// All results collected successfully
			case <-time.After(time.Second):
				t.Fatal("test timed out waiting for results")
			}
		})
	}
}

func TestWorkerPool_Cancellation(t *testing.T) {
	pool := NewWorkerPool(2)
	tasks := make(chan Task, 10)
	stop := make(chan struct{})
	results := pool.Handle(stop, tasks)

	// Send long-running tasks
	tasks <- Task{id: 1, duration: time.Second}
	tasks <- Task{id: 2, duration: time.Second}
	close(tasks)

	// Wait for workers to pick up tasks
	time.Sleep(100 * time.Millisecond)

	// Signal stop
	close(stop)

	// Create a done channel to signal when all results are collected
	done := make(chan struct{})
	go func() {
		defer close(done)
		cancelCount := 0

		for result := range results {
			if result.err != nil {
				cancelCount++
			}
		}

		if cancelCount == 0 {
			t.Error("expected cancelled tasks, got none")
		}
	}()

	// Wait for results with timeout
	select {
	case <-done:
		// All results collected successfully
	case <-time.After(2 * time.Second):
		t.Fatal("test timed out waiting for results")
	}
}
