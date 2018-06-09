package main

import (
	"database/sql"
	"log"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/gorilla/mux"
)

//Products struct
type Products struct {
	Sku          string    `form:"sku" json:"sku"`
	Product_name string    `form:"product_name" json:"product_name"`
	Stocks       int       `form:"stocks" json:"stocks"`
}

//Response struct
type ResponseProduct struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data   []Products

}

func returnAllProducts(w http.ResponseWriter, r *http.Request){
	var products Products
	var arr_products []Products
	var responseProd ResponseProduct
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/goforbeginner")
	defer db.Close()

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

	responseProd.Status = 1
	responseProd.Message = "Success"
	responseProd.Data = arr_products

	json.NewEncoder(w).Encode(responseProd)

}

func main(){

	router := mux.NewRouter()
	router.HandleFunc("/getproducts",returnAllProducts).Methods("GET")
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe(":1234",router))

}
