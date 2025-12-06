package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generators : basically, you get to generate value as and when you require

func generator[T any, K any](done <-chan K, repeatFn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)

		for {
			select {
			case <-done:
				return
			case stream <- repeatFn():
			}
		}
	}()
	return stream
}

func main() {
	randomInt := func() int {
		return rand.Intn(1000000)
	}

	done := make(chan struct{})
	randomGenerator := generator(done, randomInt)
	go func() {
		for r := range randomGenerator {
			fmt.Println(r)
		}
	}()
	time.Sleep(time.Second)
}
