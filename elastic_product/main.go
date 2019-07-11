package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/indroct/go-for-beginner/elastic_product/routers"
	"time"
	"oropay/lib/helper"
	_ "github.com/go-sql-driver/mysql" // import your required driver
	
	"github.com/astaxie/beego"
)

func init() {
	// Set database
	setDatabase()
	
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.EnableGzip = true
}

func main() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.Run()
}

func setDatabase() {
	// set default database
	configDb := beego.AppConfig.String("mysql_user") +":"+ beego.AppConfig.String("mysql_pass") +"@tcp("+
		beego.AppConfig.String("mysql_url") +")/"+ beego.AppConfig.String("mysql_db") +"?charset=utf8&parseTime=true&loc=Asia%2fJakarta"
	
	orm.RegisterDriver("mysql", orm.DRMySQL)
	
	// GMT +7
	orm.DefaultTimeLoc = helper.DefaultLocation
	
	// MaxIdleConns = 0, SetMaxOpenConns = 2400
	orm.RegisterDataBase("default", "mysql", configDb, 0, 2400)
	orm.SetDataBaseTZ("default", helper.DefaultLocation)
	
	// Connection Setting
	db, err := orm.GetDB("default")
	if err == nil {
		db.SetConnMaxLifetime(time.Second * 10)
	} else {
		beego.Error(err)
	}
	
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
}
