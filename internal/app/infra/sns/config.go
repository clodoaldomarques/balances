package sns

import (
	"balances/pkg/aws"
	"balances/pkg/logger"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func NewSNSClient(ctx context.Context) *sns.Client {
	c, err := aws.NewAWSConfig(ctx)
	if err != nil {
		logger.Fatal(ctx, "error on connect to AWS", logger.Fields{"error": err})
	}
	return sns.NewFromConfig(c)
}
