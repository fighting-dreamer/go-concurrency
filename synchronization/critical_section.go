package main

import (
	"fmt"
	"sync"
	"time"
)

func WithWaitGroup_deductAmount(amount *int, deduction int, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(index)
	*amount -= deduction
}

func withWaitGroup() {
	var amount = 10000000
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go WithWaitGroup_deductAmount(&amount, (i+1)*10000, i+1, &wg)
		fmt.Println(amount)
	}
	wg.Wait()
	fmt.Println("withWaitGroup : ", amount)

}

func WithoutWaitGroup_deductAmount(amount *int, deduction int, index int) {
	fmt.Println(index)
	*amount -= deduction

}

func withoutWaitGroup() {
	var amount = 10000000
	for i := 0; i < 5; i++ {
		go WithoutWaitGroup_deductAmount(&amount, (i+1)*10000, i+1)
		fmt.Println(amount)
	}
	time.Sleep(time.Second)
	fmt.Println("withoutWaitGroup :", amount)

}

func main() {
	withoutWaitGroup()
	withWaitGroup()
}
