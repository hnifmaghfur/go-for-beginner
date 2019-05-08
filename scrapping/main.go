package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

type Logam struct {
	Gram string
	Harga string
}

func main(){
	
	doc, err := goquery.NewDocument("https://logammulia.com/id/purchase/gold")
	
	if err != nil{
		fmt.Println(err.Error())
	}
	
	var logams []Logam
	var logam Logam
	
	doc.Find(".item-1").Each(func(idx int, s *goquery.Selection) {
		re := regexp.MustCompile("[0-9]+")
		grams := re.FindAllString(s.Text(),-1)
		fmt.Println(grams)
		gram := strings.Join(grams," ")
		logam.Gram = gram
		logams = append(logams,logam)
	})
	
	count := 0;
	doc.Find(".item-2").Each(func(idx int, s *goquery.Selection) {
		
		tmp := strings.Replace(s.Text(), "Harga", "", -1)
		tmp = strings.Replace(tmp, "Rp", "", -1)
		tmp = strings.Replace(tmp, ",", "", -1)
		tmp = strings.TrimSpace(tmp)
		fmt.Println(tmp)
		
		logams[count].Harga = tmp
		
		count++
	})
	
	
	
}
