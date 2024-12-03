package sns

import (
	"balances/pkg/aws"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func NewSNSClient(ctx context.Context) *sns.Client {
	c, err := aws.NewAWSConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return sns.NewFromConfig(c)
}
