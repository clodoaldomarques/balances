package dynamodb

import (
	"balances/pkg/aws"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Connect(ctx context.Context) *dynamodb.Client {
	c, err := aws.NewAWSConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return dynamodb.NewFromConfig(c)
}
