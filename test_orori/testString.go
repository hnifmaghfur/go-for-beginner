package main

import "strings"

func main(){
	d := "Saturday"
	
	days := strings.Split(d,",")
	
	for _,day := range days{
		println(day)
	}
}