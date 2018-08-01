// @APIVersion 1.0.0
// @Title Indra beego API
// @Description This is my research about how to make documentation in beego
// @Contact omyank2007i@gmail.com
// @TermsOfServiceUrl https://indraoctama.com/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/indraoct/go-for-beginner/beego/indra-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
