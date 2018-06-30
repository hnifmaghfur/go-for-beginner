package main

import (
	"fmt"
	"time"
)

func main() {
	n := 3

	// We want to run a goroutine to multiply n by 2
	go multiplyByTwo(n)

	// We pause the program so that the `multiplyByTwo` goroutine
	// can finish and print the output before the code exits
	time.Sleep(time.Second)
}

func multiplyByTwo(num int) int {
	result := num * 2
	fmt.Println(result)
	return result
}
