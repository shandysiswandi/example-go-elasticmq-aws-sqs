package main

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	params := &sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Title": {
				DataType:    aws.String("String"),
				StringValue: aws.String("The Whistler"),
			},
			"Author": {
				DataType:    aws.String("String"),
				StringValue: aws.String("John Grisham"),
			},
			"WeeksOn": {
				DataType:    aws.String("Number"),
				StringValue: aws.String("6"),
			},
		},
		MessageBody: aws.String("message from queue"), // Required
		QueueUrl:    aws.String(url),                  // Required
	}
	resp, _ := svc.SendMessage(params)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp.String()))
}
