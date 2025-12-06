package main

import (
	"fmt"
	"sync"
)

func WaitGroupDidnotWait() {

	chars := []byte{'a', 'b', 'c'}
	charCh := make(chan byte)

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		defer wg.Done()
		for _, ch := range chars {
			select {
			case charCh <- ch:
				fmt.Println("sent ", ch)
			}
		}
	}()

	go func(c chan byte) {
		wg.Add(1)
		defer wg.Done()
		for v := range c {
			fmt.Println(v)
		}
	}(charCh)

	wg.Wait() // this Wait executed before anyone could do the Add operation
}

func waitGroupWaitButdidnotCloseChannel() {

	chars := []byte{'a', 'b', 'c'}
	charCh := make(chan byte)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {

		defer wg.Done()
		for _, ch := range chars {
			select {
			case charCh <- ch:
				fmt.Println("sent ", ch)
			}
		}

	}()

	wg.Add(1)
	go func(c chan byte) {

		defer wg.Done()
		for v := range c {
			fmt.Println(v)
		}
	}(charCh)

	wg.Wait() // this Wait executed before anyone could do the Add operation
}

func waitGroupWaitAndCloseChannel() {
	chars := []byte{'a', 'b', 'c'}
	charCh := make(chan byte)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {

		defer wg.Done()
		for _, ch := range chars {
			select {
			case charCh <- ch:
				fmt.Println("sent ", ch)
			}
		}
		close(charCh)
	}()

	wg.Add(1)
	go func(c chan byte) {

		defer wg.Done()

		// listening a closed channel is not a problem
		// loop exits after channel is closed
		for v := range c {
			fmt.Println(v)
		}
	}(charCh)

	wg.Wait() // this Wait executed before anyone could do the Add operation
}

func waitGroupWaitButdidnotCloseBufferedChannel() {

	chars := []byte{'a', 'b', 'c'}
	charCh := make(chan byte, 2)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {

		defer wg.Done()
		for _, ch := range chars {
			select {
			case charCh <- ch:
				fmt.Println("sent ", ch)
			}
		}

	}()

	wg.Add(1)
	go func(c chan byte) {

		defer wg.Done()
		for v := range c {
			fmt.Println(v)
		}
	}(charCh)

	wg.Wait() // this Wait executed before anyone could do the Add operation
}

func waitGroupWaitAndCloseBufferedChannel() {

	chars := []byte{'a', 'b', 'c'}
	charCh := make(chan byte, 2)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {

		defer wg.Done()
		for _, ch := range chars {
			select {
			case charCh <- ch:
				fmt.Println("sent ", ch)
			}
		}
		close(charCh)
	}()

	wg.Add(1)
	go func(c chan byte) {

		defer wg.Done()
		// listening on closed channel is not a problem to program;s execution.
		// the loop exits afterwards
		for v := range c {
			fmt.Println(v)
		}
		fmt.Println("channel closed, exiting loop")
	}(charCh)

	wg.Wait() // this Wait executed before anyone could do the Add operation
}

func main() {
	// WaitGroupDidnotWait()
	// waitGroupWaitButdidnotCloseChannel()
	// waitGroupWaitAndCloseChannel()
	// waitGroupWaitButdidnotCloseBufferedChannel() // still error
	waitGroupWaitAndCloseBufferedChannel()
}
