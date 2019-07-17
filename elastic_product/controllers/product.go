package controllers

import (
	"github.com/astaxie/beego"
	"github.com/indroct/go-for-beginner/elastic_product/models"
)

type ProductController struct {
	beego.Controller
}

// @router / [get]
func(this *ProductController) GetProductData(){

	products,err := models.GetProducts()
	
	if err != nil{
		this.Ctx.Output.SetStatus(401)
		this.Data["json"] = err
	}else{
		this.Ctx.Output.SetStatus(200)
		this.Data["json"] = products
	}
	this.ServeJSON()
}