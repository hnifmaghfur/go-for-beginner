package main

import "fmt"

func main(){
	var sliceInteger []int
	a := 2
	b := 3

	channel1 := make(chan int)
	channel2 := make(chan int)
	
	go MultiPly(a,b,channel1)
	
	go Addition(a,b,channel2)
	
	sliceInteger = append(sliceInteger,<-channel1)
	sliceInteger = append(sliceInteger,<-channel2)
	
	fmt.Println(sliceInteger)
	
}


func MultiPly(a int,b int,chan1 chan<- int){
	
	result := a*b
	
	chan1 <- result
}


func Addition(a int,b int,chan2 chan<- int){
	
	result := a+b
	
	chan2 <- result
}