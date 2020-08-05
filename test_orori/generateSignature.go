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

	fmt.Println(dateString)
	fmt.Println(signature)

}