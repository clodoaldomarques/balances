package sqs

import (
	"balances/pkg/aws"
	"balances/pkg/logger"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func NewSQSClient(ctx context.Context) *sqs.Client {
	cfg, err := aws.NewCustomConfig(ctx)
	if err != nil {
		c, err := cfg.Credentials.Retrieve(ctx)
		logger.Error(ctx, err.Error(), logger.Fields{
			"KeyID":    c.AccessKeyID,
			"SecretID": c.SecretAccessKey,
			"Expires":  c.Expires,
		})
	}
	return sqs.NewFromConfig(cfg)
}
