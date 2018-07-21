package controllers

import (
	"github.com/astaxie/beego"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/orm"
	"fmt"
	jwt "github.com/juusechec/jwt-beego"
	"time"
)

type LoginController struct {
	beego.Controller
}

type ResponseLogin struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Tokenstring string `json:"tokenstring"`
}

func (this *LoginController) Post(){

	var resLogin ResponseLogin
	resLogin.Status = 0
	resLogin.Message = "Login Failed"

	hasher := md5.New()
	username := this.Ctx.Input.Query("username")
	password := this.Ctx.Input.Query("password")

	hasher.Write([]byte(password))
	password_md5 := hex.EncodeToString(hasher.Sum(nil))

	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT username, password,tokenstring FROM users WHERE username = ? AND password = ?", username,password_md5).Values(&maps)


	if err == nil && num > 0 {
		resLogin.Status = 1
		resLogin.Message = "Login Success"


		libJwt := jwt.EasyToken{Username:username,
			Expires: time. Now (). Unix () + 3600 ,}
		tokenString, _ := libJwt. GetToken ()
		resLogin.Tokenstring = tokenString
	}else{
		fmt.Println(err)
		fmt.Println(num)
		fmt.Println(password_md5)
		fmt.Println(password)
		fmt.Println(username)
	}

	this.Data["json"] = resLogin
	this.ServeJSON()

}
