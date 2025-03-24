package main

import (
	"fmt"
	"sync"
)

func mergeСhannels(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Задание 6
// Напишите программу на Go, которая сливает N каналов в один.
//
// Напишите unit тесты к созданным функциям
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

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

	go func() {
		c3 <- 5
		c3 <- 6
		close(c3)
	}()

	result := mergeСhannels(c1, c2, c3)
	for v := range result {
		fmt.Println(v)
	}
}
