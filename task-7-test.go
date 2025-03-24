package main

import (
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	const count = 10
	s := make(Semaphore, count)

	if len(s) != 0 {
		t.Fatalf("Expected semaphore to be empty, got %d", len(s))
	}

	done := make(chan struct{})

	for i := 0; i < count; i++ {
		go func(n int) {
			defer s.Add(1)
			time.Sleep(10 * time.Millisecond)
			done <- struct{}{}
		}(i)
	}

	s.Delete(count)

	for i := 0; i < count; i++ {
		<-done
	}

	if len(s) != count {
		t.Fatalf("Expected semaphore to be full after completion, got %d", len(s))
	}
}

func TestSemaphoreIncrementDecrement(t *testing.T) {
	const count = 5
	s := make(Semaphore, count)

	s.Add(count)

	if len(s) != count {
		t.Fatalf("Expected semaphore to have %d items, got %d", count, len(s))
	}

	s.Delete(count)

	if len(s) != 0 {
		t.Fatalf("Expected semaphore to be empty, got %d", len(s))
	}
}
