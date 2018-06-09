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
/**
* Seperti lazimnya web API , perlu adanya format standard untuk membangun struktur data API 
* Disini saya memecah menjadi 3 bagian yakni : 
* - status (berisi code status , misal 1 : success, 0: failed,dst)
* - message (penjelasan mengenai status)
* - data (isi data yang akan di sampaikan , dalam hal ini data produk dalam bentuk slice)
*/
type ResponseProduct struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data   []Products

}
// funtion untuk memparsing data MySQL ke JSON
func returnAllProducts(w http.ResponseWriter, r *http.Request){
	var products Products // variable untuk memetakan data product yang terbagi menjadi 3 field
	var arr_products []Products // menampung variable products ke dalam bentuk slice
	var responseProd ResponseProduct //variable untuk menampung data arr_products yang nantinya akan diubah menjadi bentuk JSON
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/goforbeginner")
	defer db.Close()

	if(err != nil) {
		log.Fatal(err)
	}

	rows, err := db.Query("Select sku,product_name,stocks from products ORDER BY sku DESC")
	if err!= nil {
		log.Print(err)
	}
	//bentuk perulangan untuk me render data dari mySQL ke struct dan slice data products
	for rows.Next(){
		if err := rows.Scan(&products.Sku, &products.Product_name, &products.Stocks); err != nil {
			log.Fatal(err.Error())

		}else{
			arr_products = append(arr_products, products)
		}
	}

	responseProd.Status = 1 //mengisi valus status = 1 , dengan asumsi pasti success
	responseProd.Message = "Success" 
	responseProd.Data = arr_products // mengisi komponen Data dengan data slice arr_products
	
	//mengubah data sstruct menjadi JSON
	json.NewEncoder(w).Encode(responseProd)

}

//fungsi main diisi dengan routing dan fungsi http GET untuk menerima response dari request HTTP
// jika program ini dijalankan maka anda bisa megkases via browser/postman dengan URL : http://localhost:1234/getproducts
func main(){

	router := mux.NewRouter()
	router.HandleFunc("/getproducts",returnAllProducts).Methods("GET") // menjalurkan URL untuk dapat mengkases data JSON API product ke /getproducts
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe(":1234",router))

}
