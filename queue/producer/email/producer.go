package main

import (
	"encoding/json"
	"github.com/indroct/go-for-beginner/test_orori/helper"
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type EmailFormat struct{
	FromEmail       string  `json:"from_email"`
	FromName        string  `json:"from_name"`
	ToEmail         string  `json:"to_email"`
	Subject         string  `json:"subject"`
	Content         string  `json:"content"`
	ContentType     string  `json:"content_type"`
	Attachment      string  `json:"attachment"`
	FileName        string  `json:"file_name"`
	Bcc             string  `json:"bcc"`
}

func main(){
	var email EmailFormat
	
	conn, err := amqp.Dial("amqp://devel:devel@178.128.23.219:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	
	q, err := ch.QueueDeclare(
		"EMAS_EMAIL", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	
	for i:= 0 ; i<5 ;i++ {
		
		email.ToEmail = "omyank2007i@gmail.com"
		email.FromEmail = "indra.octama@orori.com"
		email.Content  = " halo bro :::: "+strconv.Itoa(i)
		email.FromName = "Indra Octama"
		email.Subject = "Subject "+strconv.Itoa(i)
		
		body,_ :=json.Marshal(email)
		
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/json",
				Body:        body,
				Priority:    uint8(1),
				MessageId:   strconv.Itoa(i),
				Timestamp:   helper.GetNowTime(),
			})
		failOnError(err, "Failed to publish a message")
	}
}

