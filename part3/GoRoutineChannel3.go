package main

import "fmt"

func sender(c chan int){
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	
	close(c)
}

func main(){
	c := make(chan int,3)
	
	go sender(c)
	
	fmt.Println("Length of channel %v and capacity of channel c is %v\n",len(c),cap(c))
	
	for val := range c{
		fmt.Println("Length of channel c after value '%v' read is %v\n",val,len(c))
	}
	
}