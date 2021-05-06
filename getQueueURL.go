package main

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func GetQueueURLHandler(w http.ResponseWriter, r *http.Request) {
	urlResult, _ := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(urlResult.String()))
}
