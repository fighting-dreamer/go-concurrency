package main

import "fmt"

func sliceToChannel(nums []int) chan int {
	numsCh := make(chan int)
	go func() {
		for i := 0; i < len(nums); i++ {
			numsCh <- nums[i]
		}
		close(numsCh)
	}()
	return numsCh
}

func squareNums(numsCh chan int) chan int {
	sqNumsCh := make(chan int)
	go func() {
		for num := range numsCh {
			sqNumsCh <- num * num
		}
		close(sqNumsCh)
	}()
	return sqNumsCh
}

func sumNums(nums chan int) chan int {
	res := 0
	resCh := make(chan int)
	go func() {
		for n := range nums {
			res += n
		}
		resCh <- res
	}()

	return resCh
}

func main() {
	// you want to do sum of squares
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numsCh := sliceToChannel(nums)  // start of pipeline
	squaredCh := squareNums(numsCh) // operation/transformation one
	sumValueCh := sumNums(squaredCh)
	fmt.Println("sum : ", <-sumValueCh)
}
