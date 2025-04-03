package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestPromise(t *testing.T) {
	t.Run("successful task execution", func(t *testing.T) {
		ctx := context.Background()
		expected := "test result"

		task := func() (interface{}, error) {
			return expected, nil
		}

		resultCh := Promise(ctx, task)

		select {
		case result := <-resultCh:
			if result.Err != nil {
				t.Errorf("Expected no error, got %v", result.Err)
			}
			if result.Value != expected {
				t.Errorf("Expected %v, got %v", expected, result.Value)
			}
		case <-time.After(time.Second):
			t.Error("Test timed out")
		}
	})

	t.Run("task with error", func(t *testing.T) {
		ctx := context.Background()
		expectedErr := errors.New("test error")

		task := func() (interface{}, error) {
			return nil, expectedErr
		}

		resultCh := Promise(ctx, task)

		select {
		case result := <-resultCh:
			if result.Err != expectedErr {
				t.Errorf("Expected error %v, got %v", expectedErr, result.Err)
			}
		case <-time.After(time.Second):
			t.Error("Test timed out")
		}
	})

	t.Run("context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		task := func() (interface{}, error) {
			time.Sleep(100 * time.Millisecond)
			return "should not receive this", nil
		}

		resultCh := Promise(ctx, task)
		cancel()

		select {
		case result := <-resultCh:
			if result.Err != context.Canceled {
				t.Errorf("Expected context.Canceled error, got %v", result.Err)
			}
		case <-time.After(200 * time.Millisecond):
			t.Error("Test timed out")
		}
	})

	t.Run("heavy task cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		started := make(chan struct{})
		task := func() (interface{}, error) {
			close(started)
			time.Sleep(500 * time.Millisecond) // Simulate heavy work
			return "should not receive this", nil
		}

		resultCh := Promise(ctx, task)

		// Wait for task to start then cancel
		<-started
		cancel()

		select {
		case result := <-resultCh:
			if result.Err != context.Canceled {
				t.Errorf("Expected context.Canceled error, got %v", result.Err)
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Cancellation not handled promptly")
		}
	})

	t.Run("pre-cancelled context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel before starting

		task := func() (interface{}, error) {
			t.Error("Task should not have started")
			return nil, nil
		}

		resultCh := Promise(ctx, task)

		select {
		case result := <-resultCh:
			if result.Err != context.Canceled {
				t.Errorf("Expected context.Canceled error, got %v", result.Err)
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Should return immediately for pre-cancelled context")
		}
	})

	t.Run("no listener doesn't leak", func(t *testing.T) {
		ctx := context.Background()
		done := make(chan struct{})

		task := func() (interface{}, error) {
			defer close(done)
			return nil, nil
		}

		Promise(ctx, task) // Deliberately not reading from channel

		select {
		case <-done:
			// Test passes - goroutine completed despite no listener
		case <-time.After(100 * time.Millisecond):
			t.Error("Goroutine appears to be leaked")
		}
	})
}
