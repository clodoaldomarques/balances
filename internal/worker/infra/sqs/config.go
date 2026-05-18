package sqs

import (
	"context"

	"github.com/clodoaldomarques/balances/pkg/aws"
	"github.com/clodoaldomarques/balances/pkg/logger"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func NewSQSClient(ctx context.Context) *sqs.Client {
	cfg, err := aws.NewCustomConfig(ctx)
	if err != nil {
		logger.Fatal(ctx, "falha ao carregar configuração:", logger.Fields{
			"error":      err.Error(),
			"AwsRegion":  cfg.Region,
			"AwsAddress": cfg.BaseEndpoint,
		})
	}
	return sqs.NewFromConfig(cfg)
}
