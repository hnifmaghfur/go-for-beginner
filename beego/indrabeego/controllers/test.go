package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
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
	o := orm.NewOrm()

	var maps []orm.Params
	num, err := o.Raw("SELECT sku, product_name,stocks FROM products WHERE sku = ?", "ffffff-ccc-ikik").Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps[0]["product_name"])
	}

	var resProduct ResponseProduct
	resProduct.Status = 1
	resProduct.Message = "Success"
	resProduct.Data = maps

	this.Data["json"] = resProduct
	this.ServeJSON()



}
