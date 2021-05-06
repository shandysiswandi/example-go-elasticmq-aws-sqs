package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func CreateQueueHandler(w http.ResponseWriter, r *http.Request) {
	params := &sqs.CreateQueueInput{
		QueueName: aws.String(queueName), // Required
	}
	resp, err := svc.CreateQueue(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp.String()))
}
