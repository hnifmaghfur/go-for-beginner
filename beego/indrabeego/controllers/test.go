package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	jwt "github.com/juusechec/jwt-beego"
)

type TestController struct {
	beego.Controller
}

type ProductController struct {
	beego.Controller
}

func (this *TestController) Get() {

	this.Ctx.Output.Body([]byte(" test hello world"))
}


type ResponseProduct struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data []orm.Params
}

func (this *ProductController) Get(){

	var resProduct ResponseProduct
	var maps []orm.Params
	libJWT := jwt. EasyToken {}

	tokenstring := this.Ctx.Request.Header.Get("tokenstring")

	valid , _ , _  := libJWT. ValidateToken (tokenstring)

	if(valid == true) {

		oRM := orm.NewOrm()
		num, err := oRM.Raw("SELECT sku, product_name,stocks FROM products WHERE sku = ?", "ffffff-ccc-ikik").Values(&maps)
		if err == nil && num > 0 {
			fmt.Println(maps[0]["product_name"])
		}


		resProduct.Status = 1
		resProduct.Message = "Success"
		resProduct.Data = maps
	}else{

		resProduct.Status = 401
		resProduct.Message = "Invalid Token"
		resProduct.Data = maps
	}

	this.Data["json"] = resProduct
	this.ServeJSON()



}
