package main

import (
	"fmt"
	"html"
)

func main(){
	str := "Gedung MAN Truck & BUS"
	
	str = html.EscapeString(str)
	
	fmt.Println(str)
}
