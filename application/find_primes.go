package main

import (
	"fmt"
	"math"
	"math/rand"
)

// requirement : we want to get 100 random primes

// idea is to generate random nums and check if prime and look at only first 100 of such random

func generator[T any, K any](done chan K, fn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()
	return stream
}

func checkIsPrime(num int) bool {
	if num == 2 || num == 3 {
		return true
	}

	if num&1 == 0 {
		return false
	}

	for i := 3; i < int(math.Ceil(math.Sqrt(float64(num)))); i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// check if number is prime and put the prime one in res channel
func isPrime(numCh <-chan int, resCh chan int) {
	for num := range numCh {
		if checkIsPrime(num) {
			resCh <- num
		}
	}
}

func take(ch chan int, k int) chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		for i := 0; i < k; i++ {
			res <- <-ch
		}
	}()
	return res
}

func main() {
	randomInt := func() int {
		return rand.Intn(1000000000)
	}
	done := make(chan struct{})
	gen := generator(done, randomInt)
	primesStream := make(chan int)

	for i := 0; i < 10; i++ {
		go isPrime(gen, primesStream)
	}
	resCh := take(primesStream, 100)
	for res := range resCh {
		fmt.Println(res)
	}
	done <- struct{}{}
}
