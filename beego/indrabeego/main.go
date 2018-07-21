package main

import (
	"github.com/astaxie/beego"
	"github.com/indraoct/go-for-beginner/beego/indrabeego/controllers"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"crypto/rsa"
	"log"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init(){
	//DB Config
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@/"+beego.AppConfig.String("mysqldb")+"?charset=utf8&loc=Asia%2FJakarta")
	maxidle,_ := beego.AppConfig.Int("mysqlmaxidle")
	maxconn,_ := beego.AppConfig.Int("mysqlmaxconn")
	orm.SetMaxIdleConns("default", maxidle)
	orm.SetMaxOpenConns("default", maxconn)

	//JWT settings
	signBytes,err := ioutil.ReadFile(beego.AppConfig.String("privateKey"))
	fatal(err)

	signKey,err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes,err := ioutil.ReadFile(beego.AppConfig.String("publicKey"))
	fatal(err)

	verifyKey,err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)


}

func main() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/products", &controllers.ProductController{})
	beego.Run()
}

