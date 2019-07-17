package main

import (
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/indroct/go-for-beginner/test_orori/security"
)

const (
	AkulakuAppId  = "52191494"
	AkulakuSecKey = "WJ6CQvESawQqQxn5VJ_nd3cQdm8tGzs2cYcem3uvNsI"
	AkulakuURL    = "https://testmall.akulaku.com"
	
)

type InstallmentInfo struct {
	AppId   string   `json:"appId"`
	SkuId   string   `json:"skuId"`
	Price   string   `json:"price"`
	Qty     string   `json:"qty"`
	Version string   `json:"version"`
	Sign    string   `json:"sign"`
}


func main (){
	
	var installmentInfo InstallmentInfo
	
	//example for Installment Info
	skuId   := "gold"
	price   := "600000"
	qty     := "1"
	version := ""
	
	content := skuId+price+qty+version
	signature := generateSignature(content)
	
	installmentInfo.AppId = AkulakuAppId
	installmentInfo.SkuId = skuId
	installmentInfo.Price = price
	installmentInfo.Qty   = qty
	installmentInfo.Version = version
	installmentInfo.Sign = signature
	
	req := httplib.Get(AkulakuURL+"/api/json/public/openpay/installment/info/get.do").Debug(true)
	req.Param("appId",installmentInfo.AppId)
	req.Param("skuId",installmentInfo.SkuId)
	req.Param("price",installmentInfo.Price)
	req.Param("qty",installmentInfo.Qty)
	req.Param("version",installmentInfo.Version)
	req.Param("sign",installmentInfo.Sign)
	
	res, err := req.Bytes()
	if err != nil {
		beego.Error(err)
	}
	
	fmt.Println(string(res))
}

func generateSignature (content string) string{
	content = AkulakuAppId+AkulakuSecKey+content
	signature := base64.URLEncoding.EncodeToString([]byte(security.Sha512Encrypt(content)))
	return signature
}