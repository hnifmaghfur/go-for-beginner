package main

import (
	_ "github.com/indraoct/go-for-beginner/beego/indra-api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init(){
	//DB Config
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@/"+beego.AppConfig.String("mysqldb")+"?charset=utf8&loc=Asia%2FJakarta")
	maxidle,_ := beego.AppConfig.Int("mysqlmaxidle")
	maxconn,_ := beego.AppConfig.Int("mysqlmaxconn")
	orm.SetMaxIdleConns("default", maxidle)
	orm.SetMaxOpenConns("default", maxconn)



}

func main() {
	if beego.BConfig.RunMode == beego.AppConfig.String("runmode") {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
