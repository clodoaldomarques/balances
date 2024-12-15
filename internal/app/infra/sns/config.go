package sns

import (
	"balances/pkg/aws"
	"balances/pkg/logger"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func NewSNSClient(ctx context.Context) *sns.Client {
	cfg, err := aws.NewCustomConfig(ctx)
	if err != nil {
		logger.Fatal(ctx, "falha ao carregar configuração:", logger.Fields{
			"error":       err.Error(),
			"AwsRegion":   cfg.Region,
			"AwsAddress":  cfg.BaseEndpoint,
			"Credentials": cfg.Credentials,
		})
	}
	return sns.NewFromConfig(cfg)
}
