package sns

import (
	"context"

	"github.com/clodoaldomarques/balances/pkg/aws"
	"github.com/clodoaldomarques/balances/pkg/logger"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func NewSNSClient(ctx context.Context) *sns.Client {
	cfg, err := aws.NewCustomConfig(ctx)
	if err != nil {
		logger.Fatal(ctx, "falha ao carregar configuração:", logger.Fields{
			"error":      err.Error(),
			"AwsRegion":  cfg.Region,
			"AwsAddress": cfg.BaseEndpoint,
		})
	}
	return sns.NewFromConfig(cfg)
}
