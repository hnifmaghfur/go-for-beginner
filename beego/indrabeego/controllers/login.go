package controllers

import (
	"github.com/astaxie/beego"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/orm"
	"fmt"
	jwt "github.com/juusechec/jwt-beego"
	"time"
	"strconv"
)

type LoginController struct {
	beego.Controller
}

type ResponseLogin struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Tokenstring string `json:"tokenstring"`
	Tokenexpired string `json:"tokenexpired"`
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

	oRM := orm.NewOrm()
	var maps []orm.Params
	num, err := oRM.Raw("SELECT username, password FROM users WHERE username = ? AND password = ?", username,password_md5).Values(&maps)


	if err == nil && num > 0 {
		resLogin.Status = 1
		resLogin.Message = "Login Success"

		expires_int64 := time. Now (). Unix () + 3600
		expires := strconv.FormatInt(expires_int64,10)

		libJwt := jwt.EasyToken{Username:username,
			Expires: expires_int64 ,}
		tokenString, _ := libJwt. GetToken ()

		resLogin.Tokenstring = tokenString
		resLogin.Tokenexpired = expires
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
