package main

import (
	"e-document-api/lib/constanta"
	"encoding/json"
	"github.com/indroct/go-for-beginner/test_orori/helper"
	"log"
	"github.com/streadway/amqp")

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type PartnersLog struct{
	RequestHeader	string		`json:"request_header"`
	RequestBody		string	    `json:"request_body"`
	ResponseHeader  string      `json:"response_header"`
	ResponseBody	string		`json:"response_body"`
	ClientIp		string		`json:"client_ip"`
	PaymentType     string      `json:"payment_type"`
	Url			    string		`json:"url"`
	Method			string		`json:"method"`
	Created			string 		`json:"created"`
}

func main(){
	var logElastic PartnersLog
	
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	
	q, err := ch.QueueDeclare(
		"Log Partner", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	
	logElastic.RequestBody = "Halo test"
	logElastic.ResponseHeader = "test header"
	logElastic.ClientIp = "localhost"
	logElastic.Created = helper.GetNowTime().Format(constanta.MYSQL_DATETIME_FORMAT)
	logElastic.ResponseBody = "response body"
	logElastic.ResponseHeader = "response header"
	logElastic.Method = "POST"
	
	body,_ :=json.Marshal(logElastic)
	
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/json",
			Body:        body,
			Priority: uint8(1),
			MessageId:"Hai",
			Timestamp:helper.GetNowTime(),
		})
	failOnError(err, "Failed to publish a message")
}

