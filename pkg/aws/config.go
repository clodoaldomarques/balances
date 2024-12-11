package aws

import (
	"balances/configs"
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAWSConfig(ctx context.Context) (*session.Session, error) {
	c := configs.New()
	return session.NewSession(&aws.Config{
		Region:      aws.String(c.AwsRegion),
		Credentials: credentials.NewSharedCredentials("", c.AwsProfile),
		Endpoint:    aws.String(c.AwsAddress),
	})
}
