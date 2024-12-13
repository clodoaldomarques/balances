package sqs

import (
	"balances/pkg/aws"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func NewSQSClient(ctx context.Context) *sqs.Client {
	cfg, err := aws.NewAWSConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return sqs.NewFromConfig(cfg)
}
