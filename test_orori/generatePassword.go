package main

import (
	"fmt"
	"go-for-beginner/test_orori/security"
	"strings"
)

func main(){

	email := "indra.octama@smartfren.com"

	splitYeah := strings.Split(email,"@")
	username := splitYeah[0]
	domain   := splitYeah[1]
	hash 	 := "@Mysf2018"
	password := "skywalker"
	md5Password := security.Md5Encrypt(password)

	hashedPassword := security.Md5Encrypt(hash+username+md5Password+domain)

	fmt.Println(hashedPassword)


}

