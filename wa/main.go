package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
)

func main(){
	
	token := "GGjPKYSEfYtN8eVX67XEU4qRYiJP5hnxNTlyMAHMyjJyiKzrSaEC4ElJfCSYIcej"
	
	req := httplib.Post("https://wablas.com/api/send-message").Debug(true)
	req.Header("Authorization",token)
	req.Param("phone","082111833436")
	req.Param("message","Helo world")
	
	resp, err := req.Bytes()
	if err != nil {
		fmt.Println(err.Error())
	}
	
	fmt.Println(string(resp))
}