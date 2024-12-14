package aws

import (
	"balances/configs"
	"balances/pkg/logger"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func NewCustomConfig(ctx context.Context) (aws.Config, error) {
	env := configs.New()
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(env.AwsRegion),
		config.WithBaseEndpoint(env.AwsAddress),
	)
	if err != nil {
		logger.Fatal(ctx, err.Error(), logger.Fields{
			"region":        env.AwsRegion,
			"base_endpoing": env.AwsAddress,
		})
	}

	return cfg, nil
}
