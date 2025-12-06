package main

import "fmt"

func scenario1() {

	myCh := make(chan int)

	go func() {
		myCh <- 100
	}()

	// the myCh waits till it gets some value form some go routine passed to it.
	// you forked away to a go routine, ans you joined back.
	// refer (at this time in vedio) : https://youtu.be/qyM8Pi1KiiM?t=827
	x := <-myCh
	fmt.Println(x)
}

func scenario2() {
	myCh := make(chan int)

	go func() {
		myCh <- 100
	}()

	// the myCh waits till it gets some value form some go routine passed to it.
	// what if we dont have reciever for myCh, will the program error ?
	// : NO in this case as we main go routine completes.
	// : yes if hte main go routine waits.
	// x := <-myCh
	// fmt.Println(x)
	select {} // making indefinite wait, it will make program panic as the go-routine is stuck at myCh not able to recieve data.
}

func main() {
	// scenario1()
	scenario2()
}
