package controllers

import (
	"github.com/astaxie/beego"
	"bytes"
)

type BaseController struct {
	beego.Controller
}

func (this BaseController) Prepare (){

	if beego.BConfig.RunMode == "dev" {
		buf := new(bytes.Buffer)
		buf.ReadFrom(this.Ctx.Request.Body)
		newStr := buf.String()

		beego.Debug(this.Ctx.Request.Header)
		beego.Debug(this.Ctx.Request.Body)
		beego.Debug(newStr)
	}

}
