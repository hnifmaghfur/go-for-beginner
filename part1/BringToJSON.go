package main

import (
	"github.com/gorilla/mux"
        "log"
	"net/http"
	"fmt"
	"encoding/json"
)

type Article struct {
	Id      int    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Response struct {
	Values []Article
	Code   int
}

type Articles []Article


func returnAllArticles(w http.ResponseWriter, r *http.Request){
	var response Response
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	response.Code = 1
	response.Values = articles
	json.NewEncoder(w).Encode(response)
}

func main(){

	router := mux.NewRouter()
	router.HandleFunc("/getarticles",returnAllArticles).Methods("GET")
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe(":1234",router))

}
