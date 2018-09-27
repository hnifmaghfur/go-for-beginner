package main

import (
"fmt"
"time"
)

const (
	// See http://golang.org/pkg/time/#Parse
	timeFormat = "2006-01-02 15:04 MST"
)

func GetDiffBeforeNow(timeCheck time.Time) (timeHours float64, timeMinutes float64, timeSeconds float64){
	duration := time.Since(timeCheck)
	return float64(duration.Hours()), float64(duration.Minutes()), float64(duration.Seconds())

}

func main() {

	vd := time.Now()
	vdbanget := vd.Format(timeFormat)
	timeCheck,_ := time.Parse(timeFormat,vdbanget)

	fmt.Println(timeCheck)

	hours, minutes, seconds := GetDiffBeforeNow(timeCheck)
	fmt.Println(hours)
	fmt.Println(minutes)
	fmt.Println(seconds)
}
