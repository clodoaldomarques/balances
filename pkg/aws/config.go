package aws

import (
	"balances/configs"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func NewAWSConfig(ctx context.Context) (aws.Config, error) {
	c := configs.New()

	customEndpointResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           c.AwsAddress,
			SigningRegion: region,
		}, nil
	})

	return config.LoadDefaultConfig(
		ctx,
		config.WithRegion(c.AwsRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
		config.WithEndpointResolverWithOptions(customEndpointResolver),
	)

}
