package main

import (
	"fmt"
	"time"
)

func main(){
	done := make(chan bool, 2)
	
	go func() {
		fmt.Println("Series 1")
		time.Sleep(2 * time.Second)
		fmt.Println("Done Go Routine")
		done <- true
		done <- true
		done <- true
	}()
	
	fmt.Println("Series 2")
	<-done
	fmt.Println("Done the whole apps")
}
