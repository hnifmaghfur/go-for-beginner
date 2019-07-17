package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/indroct/go-for-beginner/elastic_product/controllers:ProductController"] = append(beego.GlobalControllerRouter["github.com/indroct/go-for-beginner/elastic_product/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetProductData",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
