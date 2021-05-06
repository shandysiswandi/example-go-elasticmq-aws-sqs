package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func connectSQSLocal() {
	// acc := os.Getenv("AWS_ELASTIQMQ_ACCESS_KEY")
	// secAcc := os.Getenv("AWS_ELASTIQMQ__SECRET_ACCESS_KEY")
	endPoint := os.Getenv("AWS_ELASTIQMQ_ENDPOINT")
	region := os.Getenv("AWS_ELASTIQMQ_REGION")
	qName := "go-queue"

	cfgSess := &aws.Config{}
	sess, _ := session.NewSession(cfgSess)

	cfgSQS := &aws.Config{Endpoint: aws.String(endPoint), Region: aws.String(region)}
	svc = sqs.New(sess, cfgSQS)
	url = fmt.Sprintf("%s/queue/%s", endPoint, qName)
	queueName = qName
}
