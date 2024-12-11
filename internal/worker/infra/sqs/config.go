package sqs

import (
	"balances/pkg/aws"
	"context"
	"log"

	"github.com/aws/aws-sdk-go/service/sqs"
)

func NewSQSClient(ctx context.Context) *sqs.SQS {
	sess, err := aws.NewAWSConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return sqs.New(sess)
}
