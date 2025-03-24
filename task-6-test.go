package main

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {
	// Тестовый случай с двумя каналами
	c1 := make(chan int)
	c2 := make(chan int)
	merged := mergeСhannels(c1, c2)

	go func() {
		c1 <- 1
		c1 <- 2
		close(c1)
	}()

	go func() {
		c2 <- 3
		c2 <- 4
		close(c2)
	}()

	values := []int{}
	for val := range merged {
		values = append(values, val)
	}

	expected := []int{1, 2, 3, 4}
	if len(values) != len(expected) {
		t.Fatalf("Expected %d values, got %d", len(expected), len(values))
	}

	for i, v := range expected {
		if values[i] != v {
			t.Fatalf("Expected %d, got %d at index %d", v, values[i], i)
		}
	}
}

func TestMergeEmptyChannels(t *testing.T) {
	merged := mergeСhannels()

	values := []int{}
	for val := range merged {
		values = append(values, val)
	}

	if len(values) != 0 {
		t.Fatalf("Expected no values, got %d", len(values))
	}
}

func TestMergeSingleChannel(t *testing.T) {
	ch := make(chan int)
	merged := mergeСhannels()

	go func() {
		ch <- 5
		ch <- 10
		close(ch)
	}()

	values := []int{}
	for val := range merged {
		values = append(values, val)
	}

	expected := []int{5, 10}
	if len(values) != len(expected) {
		t.Fatalf("Expected %d values, got %d", len(expected), len(values))
	}

	for i, v := range expected {
		if values[i] != v {
			t.Fatalf("Expected %d, got %d at index %d", v, values[i], i)
		}
	}
}
