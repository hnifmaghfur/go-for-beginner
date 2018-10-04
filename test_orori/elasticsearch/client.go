package main

import (
	"github.com/olivere/elastic"
	"context"
	"time"
	"os"
	"log"
	"fmt"
)

type tweet struct {
	User     string        `json:"user"`
	Message  string        `json:"message"`
	Retweets int           `json:"retweets"`
	Image    string        `json:"image,omitempty"`
	Created  string     `json:"created,omitempty"`
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
	rs,err :=createTwitterIndex(client,"twitter")
	if(err != nil){
		fmt.Println("ERROR: ",err.Error())
	}
	fmt.Println("RESPONSE: ",rs)

	insertData(client,"twitter")

}

func createTwitterIndex(client *elastic.Client,index string)(rs string, err error){

	if(isExistIndexTwitter(client,index) == false) {
		// Create a new index.
		mapping := `{
			"settings":{
				"number_of_shards":1,
				"number_of_replicas":0
			},
			"mappings":{
				"tweet":{
					"properties":{
						"tags":{
							"type":"text"
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
		createIndex, err := client.CreateIndex("twitter").BodyString(mapping).Do(ctx)
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

func isExistIndexTwitter(client *elastic.Client,index string) (b bool){
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
		Id("1").
		BodyJson(tweet1).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

}