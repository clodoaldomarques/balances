package sns

import (
	"balances/pkg/aws"
	"balances/pkg/logger"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func NewSNSClient(ctx context.Context) *sns.Client {
	cfg, err := aws.NewAWSConfig(ctx)
	if err != nil {
		logger.Fatal(ctx, err.Error(), logger.Fields{})
	}
	return sns.NewFromConfig(cfg)
}
