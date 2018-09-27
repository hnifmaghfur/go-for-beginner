package main

import "time"
import "github.com/indraoct/go-for-beginner/test_orori/security"
import (
	"github.com/indraoct/go-for-beginner/test_orori/helper"
	"fmt"
)

type OauthClient struct {
	Id       		uint64			`json:"id"`
	Name 			string			`orm:"size(100)" valid:"Required" json:"name"`
	SecretKey		string			`orm:"size(100)" valid:"Required" json:"secret_key"`
	TokenExpires	uint32			`valid:"Required" json:"token_expires"`
	Email			string			`valid:"Email" orm:"size(100)" json:"email"`
	IsActive		uint8			`orm:"default(1)" json:"is_active"`
	CreatedAt		time.Time 		`orm:"auto_now_add;type(datetime)" json:"created_at"`
	CreatedBy		string	  		`json:"created_by"`
	UpdatedAt		time.Time 		`orm:"auto_now;type(datetime)"`
	UpdatedBy		string 			`orm:"null"`
}

type User struct{
	Id       		uint64		`json:"id"`
	GroupId       	uint64		`json:"group_id"`
	Username		string		`orm:"size(50)" valid:"Required" json:"username"`
	Password 		string		`orm:"size(100)" valid:"Required" json:"password"`
	Salt  	 		string		`orm:"size(100)" valid:"Required" json:"salt"`
	LoginNumber		int64		`json:"login_number"`
	Token			string		`orm:"size(100)" json:"token"`
	LastLogin		time.Time	`orm:"type(datetime)" json:"last_login"`
	IsLogin			uint8		`orm:"default(1)" json:"is_login"`
	IsActive		uint8		`orm:"default(1)" json:"is_active"`
	CreatedAt		time.Time 	`orm:"auto_now_add;type(datetime)" json:"created_at"`
	CreatedBy		string	  	`json:"created_by"`
	UpdatedAt		time.Time 	`orm:"auto_now;type(datetime)" json:"updated_at"`
	UpdatedBy		string 		`orm:"null" json:"updated_by"`
	OauthClient     *OauthClient	`orm:"rel(fk);column(oauth_client_id)" json:"oauth_client"`
}

func main(){

	var usr User
	usr.Username = "omyank20071@gmail.com"
	usr.Password = "123456"

	createuser(usr)

	loginuser(usr)

	generateCustomerSign()

}

func createuser(usr User){
	usr.Salt = security.ShaOneEncrypt(helper.GetNowTime().String() + helper.StringRandomWithCharset(32, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))
	usr.Password = security.ShaOneEncrypt(security.Md5Encrypt(usr.Password + usr.Salt))

	fmt.Println("-------CREATE USER--------")
	fmt.Println(usr.Salt)
	fmt.Println(usr.Password)

}

func generateCustomerSign(){
	appId := "TOKOPEDIA"
	timeStamp := "2017-12-01 00:00:00"
	SecretKey := "5affc2f9494cd32b0da015dc3ddccc0ba62a9486"
	sign := security.ShaOneEncrypt(security.Md5Encrypt(appId + timeStamp + SecretKey))

	fmt.Println("---- Generate Customer Signature ----")
	fmt.Println(sign)
}


func loginuser(usr User){

	usr.Token = security.ShaOneEncrypt(security.ShaOneEncrypt(helper.GetNowTime().String()))

	fmt.Println("-------LOGIN USER--------")
	fmt.Println(usr.Token)
}
