package main

import (
	"fmt"
	"go-for-beginner/test_orori/constanta"
	"go-for-beginner/test_orori/helper"
	"go-for-beginner/test_orori/security"
)

func main(){
	fmt.Println("FIRST ATTEMP:")
    fmt.Println(helper.GetNowTime().Unix())
	key := "MDRjMjJiNzk1MjFmNGIzYjk2OTIzYzI5YWZiNGM5YWE="
	//dateString := helper.GetNowTime().Format(constanta.MST_DATETIME_FORMAT)

	dateString := helper.GetNowTime().Format(constanta.MST_DATETIME_FORMAT)
	signature := security.Base64Encode(security.Sha256_HMAC(key,dateString))

	fmt.Println("dateString",dateString)
	fmt.Println("signature",signature)
	fmt.Println("authorization","Signature keyid=5d1d334b00d0463f7934d511952e4754f4b14b2a8c6be22d9ec83e03,algorithm=hmac-sha256,signature="+signature)

}