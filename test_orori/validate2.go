package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("2006-01-02 15:04:09", "2011-01-19 22:15:09")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.String())
}