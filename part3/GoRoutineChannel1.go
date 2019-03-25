package main

import (
	"fmt"
)

func main() {
	n := 3

	// This is where we "make" the channel, which can be used
	// to move the `int` datatype
	out := make(chan int)
	out2 := make(chan int)

	// We still run this function as a goroutine, but this time,
	// the channel that we made is also provided
	
	go multiplyByThree(n, out)
	
	go multiplyByTwo(n, out2)

	// Once any output is received on this channel, print it to the console and proceed
	fmt.Println(<-out)
	fmt.Println(<-out2)
}

// This function now accepts a channel as its second argument...
func multiplyByTwo(num int, out chan<- int) {
	result := num * 2

	//... and pipes the result into it
	out <- result
}

func multiplyByThree(num int, out chan<- int){
	result := num * 3
	
	out <- result
}
