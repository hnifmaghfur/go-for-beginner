package main

import (
	"github.com/indraoct/go-for-beginner/test_orori/security"
	"fmt"
)


func main(){
	timestamp :="123"
	sign := security.ShaOneEncrypt(security.Md5Encrypt("GOLDRUSH" + timestamp + "661ee457759b2b58bdd78059601fc5591f9c7fe5"))

	fmt.Println(sign)
}
