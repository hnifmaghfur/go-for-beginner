package main

import (
"crypto/hmac"
"crypto/sha256"
"encoding/base64"
"fmt"
)

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func main() {
	fmt.Println(ComputeHmac256("POST:/pay/payrequest:2aa8eae2-9408-43fe-aaa4-922edff6ab25:2019-04-12T11:29:00.017+07:00:{\"txn_amount\":\"154250000\",\"cashtag\":\"$getha36\",\"promo_code\":\"SERBA60\",\"url_callback\":\"https://e-mas.com\",\"purchase_desc\":\"PEMBELIANdiTOKO123.com\"}", "3800092a-9fbb-40ad-90ca-02808a44ed06"))
}