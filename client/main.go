package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joho/godotenv"
)

var svc *sqs.SQS
var url string

func main() {
	// load env
	if err := godotenv.Load(); err != nil {
		println(err.Error())
		return
	}

	// connection
	endPoint := os.Getenv("AWS_ELASTIQMQ_ENDPOINT")
	region := os.Getenv("AWS_ELASTIQMQ_REGION")
	qName := "go-queue"

	cfgSess := &aws.Config{}
	sess, _ := session.NewSession(cfgSess)

	cfgSQS := &aws.Config{Endpoint: aws.String(endPoint), Region: aws.String(region)}
	svc = sqs.New(sess, cfgSQS)
	url = fmt.Sprintf("%s/queue/%s", endPoint, qName)

	// receive queue
	for {
		time.Sleep(time.Millisecond * 500)
		receiveMessage()
		break
	}
}

func receiveMessage() {
	params := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(url), // Required
	}
	resp, err := svc.ReceiveMessage(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
}
