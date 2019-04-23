package main

import (
	"encoding/json"
	"fmt"
	"github.com/indroct/go-for-beginner/test_orori/security"
	"strings"
)

type PayRequest struct {
	TxnAmount       string  `json:"txn_amount"`
	Cashtag         string  `json:"cashtag"`
	PromoCode       string  `json:"promo_code"`
	UrlCallback     string  `json:"url_callback"`
	PurchaseDesc    string  `json:"purchase_desc"`
}
//
//{
//"txn_amount":"154250000",
//"cashtag":"$jenpay36",
//"promo_code":"SERBA60",
//"url_callback":"http://callback.url/notify",
//"purchase_desc":"PEMBELIAN di TOKO123.com"
//}

func main(){
	var rq PayRequest
	
	rq.TxnAmount = "154250000"
	rq.Cashtag = "$getha36"
	rq.PromoCode = "SERBA60"
	rq.UrlCallback = "https://e-mas.com"
	rq.PurchaseDesc = "PEMBELIAN di TOKO123.com"
	
	requestBody,_:=json.Marshal(rq)
	
	testJson := string(requestBody)
	
	
	
	fmt.Println(testJson)
	
	client_id := "e0e131c2-c2de-41bb-a6e9-edff8c9731d0:f3db9140-a451-4519-9217-f3d0f9f1c851"
	
	client_id_base64 := security.Base64Encode(client_id)
	
	if client_id_base64 == "ZTBlMTMxYzItYzJkZS00MWJiLWE2ZTktZWRmZjhjOTczMWQwOmYzZGI5MTQwLWE0NTEtNDUxOS05MjE3LWYzZDBmOWYxYzg1MQ=="{
		fmt.Println("bener mas bro")
	}else{
		fmt.Println("salah dab!")
	}
	
	apiKey    := "2aa8eae2-9408-43fe-aaa4-922edff6ab25"
	apiSecret := "3800092a-9fbb-40ad-90ca-02808a44ed06"
	Timestamp := "2019-04-12T11:29:00.017+07:00"
	
	
	httpVerb := "POST"
	relativeUrl := "/pay/payrequest"
	
	StringtoSignature :=httpVerb + ":" + relativeUrl + ":" + apiKey + ":" +
		Timestamp + ":" + testJson
	
	StringtoSignature = strings.Replace(StringtoSignature," ","",-1)
	
	
	BTPNSignature := security.ComputeHmac256(StringtoSignature,apiSecret)
	
	fmt.Println("string_to_signature",StringtoSignature)
	
	fmt.Println("BTPN-Signature: "+BTPNSignature)
	
	
	
	
}