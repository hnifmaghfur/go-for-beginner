package main

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//Products struct
type Products struct {
	Sku          string    `form:"sku" json:"sku"`
	Product_name string    `form:"product_name" json:"product_name"`
	Stocks       int       `form:"stocks" json:"stocks"`
}

func main(){
	var products Products
	var arr_products []Products
	db, err := sql.Open("mysql","root:blink182@tcp(127.0.0.1:3306)/tokoijah")

	if(err != nil) {
		log.Fatal(err)
	}
	rows, err := db.Query("Select sku,product_name,stocks from products ORDER BY sku DESC")
	if err!= nil {
		log.Print(err)
	}
	for rows.Next(){
		if err := rows.Scan(&products.Sku, &products.Product_name, &products.Stocks); err != nil {
			log.Fatal(err.Error())

		}else{
			arr_products = append(arr_products, products)
		}
	}

	for i := 0; i < len(arr_products); i++ {

		fmt.Println(arr_products[i])
	}


}
