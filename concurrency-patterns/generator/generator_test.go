package main

import (
	"context"
	"testing"
	"time"
)

func TestGenerator(t *testing.T) {
	t.Run("successful generation", func(t *testing.T) {
		ctx := context.Background()
		data := []interface{}{1, "two", 3.0, true}

		ch := Generator(ctx, data)

		for i, expected := range data {
			select {
			case val := <-ch:
				if val != expected {
					t.Errorf("item %d: expected %v, got %v", i, expected, val)
				}
			case <-time.After(100 * time.Millisecond):
				t.Errorf("timeout waiting for item %d", i)
			}
		}

		// Verify channel is closed
		if val, ok := <-ch; ok {
			t.Errorf("channel should be closed but received: %v", val)
		}
	})

	t.Run("empty data", func(t *testing.T) {
		ctx := context.Background()
		ch := Generator(ctx, []interface{}{})

		select {
		case val, ok := <-ch:
			if ok {
				t.Errorf("expected closed channel, got value: %v", val)
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("timeout waiting for channel to close")
		}
	})

	t.Run("context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Create large dataset to ensure we can cancel mid-generation
		data := make([]interface{}, 1000)
		for i := range data {
			data[i] = i
		}

		ch := Generator(ctx, data)

		// Read a few items
		for i := 0; i < 3; i++ {
			select {
			case val := <-ch:
				if val != i {
					t.Errorf("item %d: expected %v, got %v", i, i, val)
				}
			case <-time.After(200 * time.Millisecond):
				t.Errorf("timeout waiting for item %d", i)
			}
		}

		// Cancel context
		cancel()

		// Verify channel is closed after cancellation
		select {
		case val, ok := <-ch:
			if ok {
				t.Errorf("channel should be closed but received: %v", val)
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("timeout waiting for channel to close after cancellation")
		}
	})

	t.Run("pre-cancelled context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel before starting

		data := []interface{}{1, 2, 3}
		ch := Generator(ctx, data)

		// Verify channel is immediately closed
		select {
		case val, ok := <-ch:
			if ok {
				t.Errorf("channel should be closed but received: %v", val)
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("timeout waiting for channel to close")
		}
	})

	t.Run("nil data", func(t *testing.T) {
		ctx := context.Background()
		ch := Generator(ctx, nil)

		select {
		case val, ok := <-ch:
			if ok {
				t.Errorf("expected closed channel, got value: %v", val)
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("timeout waiting for channel to close")
		}
	})

	t.Run("immediate cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		ch := Generator(ctx, []interface{}{1, 2, 3})
		select {
		case _, ok := <-ch:
			if ok {
				t.Error("channel should be closed")
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("timeout waiting for channel to close")
		}
	})
}
