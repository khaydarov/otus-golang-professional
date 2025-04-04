package main

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestWorker_Handle(t *testing.T) {
	tests := []struct {
		name     string
		tasks    []Task
		wantErr  bool
		cancelAt time.Duration
	}{
		{
			name: "single task",
			tasks: []Task{
				{id: 1, duration: 10 * time.Millisecond},
			},
		},
		{
			name: "multiple tasks",
			tasks: []Task{
				{id: 1, duration: 10 * time.Millisecond},
				{id: 2, duration: 20 * time.Millisecond},
			},
		},
		{
			name: "cancel during task",
			tasks: []Task{
				{id: 1, duration: 100 * time.Millisecond},
			},
			wantErr:  true,
			cancelAt: 50 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Worker{id: 1, lastTaskTime: time.Now()}
			tasks := make(chan Task, len(tt.tasks))
			results := make(chan TaskResult, len(tt.tasks))
			var wg sync.WaitGroup
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// Start worker
			wg.Add(1)
			go w.Handle(ctx, tasks, results, &wg)

			// Send tasks
			for _, task := range tt.tasks {
				tasks <- task
			}
			close(tasks)

			if tt.cancelAt > 0 {
				time.AfterFunc(tt.cancelAt, cancel)
			}

			// Create a done channel to signal when all results are collected
			done := make(chan struct{})
			go func() {
				defer close(done)
				count := 0
				errCount := 0

				// Wait for worker to finish
				wg.Wait()
				close(results)

				// Collect all results
				for result := range results {
					count++
					if result.err != nil {
						errCount++
					}
				}

				if tt.wantErr && errCount == 0 {
					t.Error("expected errors but got none")
				}

				if !tt.wantErr && errCount > 0 {
					t.Errorf("expected no errors but got %d", errCount)
				}

				// Check metrics
				if w.tasksProcessed != count-errCount {
					t.Errorf("processed tasks mismatch: got %d, want %d", w.tasksProcessed, count-errCount)
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

func TestWorker_Metrics(t *testing.T) {
	w := &Worker{id: 1, lastTaskTime: time.Now()}
	w.tasksProcessed = 5
	w.totalWorkTime = time.Second * 10

	metrics := w.GetMetrics()
	expected := "Worker 1: 5 tasks processed, 10s total work time"
	if metrics != expected {
		t.Errorf("GetMetrics() = %v, want %v", metrics, expected)
	}
}
