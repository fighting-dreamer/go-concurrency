package main

import (
	"fmt"
	"time"
)

func indefinite_running_goroutine() {
	go func() {
		for {
			select {
			default:
				fmt.Println("work done : indefinite_running_goroutine")
			}
		}
	}()

	time.Sleep(time.Second * 10)
}

func definite_running_goroutine_using_donechannel() {
	done := make(chan struct{})
	go func(d chan struct{}) {
		for {
			select {
			case <-d:
				fmt.Println("work is complete : definite_running_goroutine_using_donechannel")
				return
			default:
				fmt.Println("work done : definite_running_goroutine_using_donechannel")
			}
		}
	}(done)
	time.Sleep(time.Second * 1)
	done <- struct{}{}
}

func main() {
	// indefinite_running_goroutine()
	definite_running_goroutine_using_donechannel()
}
