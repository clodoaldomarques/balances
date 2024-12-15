package aws

import (
	"balances/configs"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func NewCustomCredentials(c *configs.Config) aws.CredentialsProvider {
	return aws.NewCredentialsCache(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		creds := aws.Credentials{
			AccessKeyID:     c.AccessKeyID,
			SecretAccessKey: c.SecretAccessKey,
			Source:          "Environment",
		}
		if creds.AccessKeyID == "" || creds.SecretAccessKey == "" {
			return aws.Credentials{}, fmt.Errorf("credenciais AWS ausentes nas variáveis de ambiente")
		}
		return creds, nil
	}))
}

func NewCustomConfig(ctx context.Context) (aws.Config, error) {
	c := configs.New()
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(c.AwsRegion),
		config.WithBaseEndpoint(c.AwsAddress),
		config.WithCredentialsProvider(NewCustomCredentials(c)),
	)

	if err != nil {
		return aws.Config{}, fmt.Errorf("falha ao carregar configuração AWS: %w", err)
	}

	return cfg, nil
}
