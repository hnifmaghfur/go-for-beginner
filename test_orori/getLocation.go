package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

//{
//"area_code": 0,
//"city": "Jakarta",
//"company": "Pt Mitra Visioner Pratama",
//"continent_code": "AS",
//"country_code": "ID",
//"country_code3": "IDN",
//"country_name": "Indonesia",
//"found": 1,
//"ip": "103.80.237.18",
//"ip_header": "Your IP address",
//"lat": -6.1744,
//"lng": 106.8294,
//"metro_code": 0,
//"postal_code": null,
//"region": "04",
//"region_name": "Jakarta Raya",
//"time_zone": "Asia/Jakarta"
//}

type GetLocation struct{
	AreaCode int `json:"area_code"`
	City string `json:"city"`
	Company string `json:"company"`
	ContinentCode string `json:"continent_code"`
	CountryCode string `json:"country_code"`
	CountryCode3 string `json:"country_code_3"`
	CountryName string `json:"country_name"`
	Found int `json:"found"`
	Ip string `json:"ip"`
	IpHeader string `json:"ip_header"`
	Lat string `json:"lat"`
	Lng string `json:"lng"`
	MetroCode string `json:"metro_code"`
	PostalCode string `json:"postal_code"`
	Region string `json:"region"`
	RegionName string `json:"region_name"`
	TimeZone string `json:"time_zone"`
}



func main() {

	var getLocation GetLocation
	req := httplib.Post("https://iplocation.com/").Debug(true)
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	req.Header("APIKey", beego.AppConfig.String("mandiri_ecash_api_key"))
	req.Param("ip", "103.80.237.18")

	req.ToJSON(&getLocation)

	fmt.Println(getLocation.Found)
	fmt.Println(getLocation.City+", "+getLocation.CountryName)
}