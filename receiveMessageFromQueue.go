package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ReceiveMessageHandler(w http.ResponseWriter, r *http.Request) {
	params := &sqs.ReceiveMessageInput{
		QueueUrl:              aws.String(url), // Required
		AttributeNames:        aws.StringSlice([]string{"All"}),
		MessageAttributeNames: aws.StringSlice([]string{"All"}), // optional (for filter)
	}
	resp, _ := svc.ReceiveMessage(params)

	temp := ""
	for _, v := range resp.Messages {
		log.Println("1", v.Attributes)
		log.Println("2", v.MessageAttributes)
		log.Println("body", *v.Body)
		log.Println("mID", *v.MessageId)
		temp = temp + " " + v.String() + ", "
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(temp))
}
