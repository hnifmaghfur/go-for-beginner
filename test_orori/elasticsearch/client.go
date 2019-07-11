package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"os"
	"time"
)

type tweet struct {
	User     string        `json:"user"`
	Message  string        `json:"message"`
	Retweets int           `json:"retweets"`
	Image    string        `json:"image,omitempty"`
	Created  string        `json:"created,omitempty"`
	Tags     []string      `json:"tags,omitempty"`
	Location string        `json:"location,omitempty"`
}


func main() {

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200", "http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))

	if(err != nil){
		panic(err)
	}
	//create twitter index
	rs,err :=createMustangIndex(client,"mustang")
	if(err != nil){
		fmt.Println("ERROR: ",err.Error())
	}
	fmt.Println("RESPONSE: ",rs)
	
	insertData(client,"mustang")
	
	ctx := context.Background()
	//get1, err := client.Get().
	//	Index("twitter").
	//	Id("1").
	//	Do(ctx)
	//
	//if err != nil{
	//	fmt.Println(err.Error())
	//}
	//
	//rawdata,_ := get1.Source.MarshalJSON()
	//
	//if get1.Found{
	//	fmt.Printf(string(rawdata))
	//}
	
	
	////search
	termQuery := elastic.NewTermQuery("user", "indra96")
	searchResult, err := client.Search().
		Index("mustang").   // search in index "twitter"
		Query(termQuery).   // specify the query
		Sort("user", true). // sort by "user" field, ascending
		From(0).Size(9).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)             // execute
	
	if err != nil {
		// Handle error
		fmt.Println(err.Error())
	}
	
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	

}

func createMustangIndex(client *elastic.Client,index string)(rs string, err error){

	if(isExistIndexMustang(client,index) == false) {
		// Create a new index.
		mapping := `{
			"settings":{
				"number_of_shards":1,
				"number_of_replicas":0
			},
			"mappings":{
				"tweet":{
					"properties":{
						"user":{
							"type":"keyword"
						},
						"message":{
							"type":"text",
							"store": true,
							"fielddata": true
						},
						"image":{
							"type":"keyword"
						},
						"created":{
							"type":"date"
						},
						"tags":{
							"type":"keyword"
						},
						"location":{
							"type":"geo_point"
						},
						"suggest_field":{
							"type":"completion"
						}
					}
				}
			}
		}`
		ctx := context.Background()
		createIndex, err := client.CreateIndex("mustang").BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			return "Not OK",err
		}
		if !createIndex.Acknowledged {
			return "Not Aknowledge",nil
		}
	}else{
		return "Index: "+index+" is exist",nil
	}

	return "OK",nil
}

func isExistIndexMustang(client *elastic.Client,index string) (b bool){
	// Check if the index called "twitter" exists
	exists, err := client.IndexExists(index).Do(context.Background())
	if err != nil {
		return false
	}
	if !exists {
		return false
	}

	return true
}

func insertData(client *elastic.Client,index string){



	// Index a tweet (using JSON serialization)
	ctx := context.Background()
	tweet1 := tweet{User: "indra96", Message: "Take Matamu", Retweets: 0, Created:time.Now().Format(time.RFC3339)}
	put1, err := client.Index().
		Index(index).
		Type("tweet").
		Id("2").
		BodyJson(tweet1).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

}

