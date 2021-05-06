package main

import (
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// source : https://gist.github.com/p4tin/f339e521556e83bee6dd
// source : http://www.inanzzz.com/index.php/post/x4cy/a-simple-aws-sqs-example-with-golang-suing-localstack

var queueName string
var svc *sqs.SQS
var url string

func main() {
	// load env
	if err := godotenv.Load(); err != nil {
		println(err.Error())
	}

	connectSQSLocal()

	r := mux.NewRouter()

	r.HandleFunc("/create-queue", CreateQueueHandler)
	r.HandleFunc("/get-queue-url", GetQueueURLHandler)
	r.HandleFunc("/list-queue-url", ListQueueURLHandler)
	r.HandleFunc("/send-message", SendMessageHandler)
	r.HandleFunc("/receive-message", ReceiveMessageHandler)

	log.Println("server run on port : 4000")
	s := &http.Server{
		Addr:           ":4000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
