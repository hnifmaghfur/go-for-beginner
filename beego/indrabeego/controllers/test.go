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

func (c *TestController) Get() {

	c.Ctx.Output.Body([]byte(" test hello world"))
}

type Products struct {
	Sku string
	Product_name string
	Stocks int
}

func (c *ProductController) Get(){
	o := orm.NewOrm()

	var maps []orm.Params
	num, err := o.Raw("SELECT sku, product_name,stocks FROM products WHERE sku = ?", "ffffff-ccc-ikik").Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps[0]["product_name"])
	}



}
