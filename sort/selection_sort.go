package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	slice := generateRandomString(20)
	fmt.Println("\n=========== Unsort Number ======\n\n",slice)
	
	selectionSort(slice)
	fmt.Println("\n========= Sorted Number =====\n\n",slice)
	
	
}

func generateRandomString(size int) []int{

	slice := make([]int,size,size)
	rand.Seed(time.Now().UnixNano())
	
	for i:= 0;i< size; i++{
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice

}

func selectionSort(items []int){
	var n = len(items)
	for i:= 0 ; i< n; i++{
		var minIdx = i
		for j := i ; j < n; j++ {
			if items[j] < items[minIdx]{
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}


