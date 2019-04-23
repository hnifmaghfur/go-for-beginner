package main

import (
	"encoding/json"
	"fmt"
	"github.com/indroct/go-for-beginner/test_orori/redis/redis"
)

func main(){
	var ttt []int
	
	ttt = append(ttt,2)
	
	bbb,_ :=json.Marshal(&ttt)
	
	redis.Set("tttt",string(bbb))
	
	rs,err :=redis.GetBytes("GET","tttt")
	
	if err != nil{
		fmt.Println(err.Error())
	}
	
	json.Unmarshal(rs,&ttt)
	
	for _,value := range ttt{
		fmt.Println(value)
	}
}
