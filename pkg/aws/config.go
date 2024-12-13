package aws

import (
	"balances/configs"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func NewAWSConfig(ctx context.Context) (aws.Config, error) {
	cfg := configs.New()
	return config.LoadDefaultConfig(
		ctx,
		config.WithSharedConfigProfile(cfg.AwsProfile),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: cfg.AwsAddress}, nil
			}),
		),
	)
}
