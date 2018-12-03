package main

import (
	"fmt"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(100)
	fmt.Printf("Threads: %v\n",runtime.GOMAXPROCS(-1))
}
