package main

import (
	"testing"
)

func TestChannelProcessing(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	go func() {
		for i := uint8(0); i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	go func() {
		defer close(out)
		for {
			num, isOpened := <-in
			if !isOpened {
				return
			}
			out <- float64(num * num * num)
		}
	}()

	expectedResults := []float64{0, 1, 8, 27, 64, 125, 216, 343, 512, 729}
	i := 0
	for num := range out {
		if num != expectedResults[i] {
			t.Errorf("Expected %f, got %f", expectedResults[i], num)
		}
		i++
	}
}
