package helper

import (
	"github.com/astaxie/beego"
	//"oropay/lib/sprint"
)

func SendSms(hp string, msg string) {
	// If Telkomsel
	//if thirdparty.TselPrefix(hp) {
	//	clickatell.SendSms(hp, msg)
	//} else {
	//	infobip.SendSms(hp, msg)
	//}

	// Only send if it's production to minimize cost
	if beego.BConfig.RunMode == "prod" {
		//sprint.SendSms(hp, msg)
	}
}