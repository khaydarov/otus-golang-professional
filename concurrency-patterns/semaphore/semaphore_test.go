package main

import (
	"sync"
	"testing"
	"time"
)

func TestNewSemaphore(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
	}{
		{"zero capacity", 0},
		{"positive capacity", 5},
		{"large capacity", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sem := NewSemaphore(tt.capacity)
			if sem == nil {
				t.Fatal("NewSemaphore returned nil")
			}
			if cap := sem.Cap(); cap != tt.capacity {
				t.Errorf("expected capacity %d, got %d", tt.capacity, cap)
			}
		})
	}
}

func TestSemaphore_Basic(t *testing.T) {
	sem := NewSemaphore(2)

	if cap := sem.Cap(); cap != 2 {
		t.Errorf("expected capacity 2, got %d", cap)
	}

	if len := sem.Len(); len != 0 {
		t.Errorf("expected length 0, got %d", len)
	}

	// Should be able to acquire twice
	sem.Acquire()
	sem.Acquire()

	if len := sem.Len(); len != 2 {
		t.Errorf("expected length 2 after acquires, got %d", len)
	}

	// Release one
	sem.Release()
	if len := sem.Len(); len != 1 {
		t.Errorf("expected length 1 after release, got %d", len)
	}
}

func TestSemaphore_TryAcquire(t *testing.T) {
	sem := NewSemaphore(1)

	// First try should succeed
	if !sem.TryAcquire() {
		t.Error("first TryAcquire should succeed")
	}

	// Second try should fail
	if sem.TryAcquire() {
		t.Error("second TryAcquire should fail")
	}

	// After release, should succeed again
	sem.Release()
	if !sem.TryAcquire() {
		t.Error("TryAcquire after release should succeed")
	}
}

func TestSemaphore_Concurrent(t *testing.T) {
	sem := NewSemaphore(3)
	var wg sync.WaitGroup
	counter := 0
	var mu sync.Mutex

	// Launch 10 goroutines trying to increment counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem.Acquire()
			defer sem.Release()

			mu.Lock()
			counter++
			time.Sleep(time.Millisecond) // Simulate work
			mu.Unlock()
		}()
	}

	wg.Wait()
	if counter != 10 {
		t.Errorf("expected counter to be 10, got %d", counter)
	}
}

func TestSemaphore_Close(t *testing.T) {
	sem := NewSemaphore(1)
	sem.Acquire()
	sem.Close()

	// After close, Release should panic
	defer func() {
		if r := recover(); r != nil {
			t.Error("Release after Close should panic")
		}
	}()
	sem.Release()
}

func TestSemaphore_Blocking(t *testing.T) {
	sem := NewSemaphore(1)
	sem.Acquire()

	// Start a goroutine that will block on acquire
	done := make(chan bool)
	go func() {
		sem.Acquire()
		done <- true
	}()

	// Should not receive from done immediately
	select {
	case <-done:
		t.Error("Acquire should block when semaphore is full")
	case <-time.After(time.Millisecond * 100):
		// This is expected
	}

	// After release, should receive from done
	sem.Release()
	select {
	case <-done:
		// This is expected
	case <-time.After(time.Millisecond * 100):
		t.Error("Acquire should succeed after Release")
	}
}

func BenchmarkSemaphore_AcquireRelease(b *testing.B) {
	sem := NewSemaphore(b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sem.Acquire()
		sem.Release()
	}
}

func BenchmarkSemaphore_TryAcquire(b *testing.B) {
	sem := NewSemaphore(b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if sem.TryAcquire() {
			sem.Release()
		}
	}
}

func BenchmarkSemaphore_Parallel(b *testing.B) {
	sem := NewSemaphore(100)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sem.Acquire()
			sem.Release()
		}
	})
}
