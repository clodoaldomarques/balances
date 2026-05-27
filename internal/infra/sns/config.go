package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/clodoaldomarques/balances-api/config"
	"github.com/clodoaldomarques/core-sdk/pkg/aws"
	"github.com/clodoaldomarques/core-sdk/pkg/logger"
)

func NewSNSClient(ctx context.Context) *sns.Client {
	cfg, err := aws.NewCustomConfig(ctx, config.New())
	if err != nil {
		logger.Fatal(ctx, "falha ao carregar configuração:", logger.Fields{
			"error":      err.Error(),
			"AwsRegion":  cfg.Region,
			"AwsAddress": cfg.BaseEndpoint,
		})
	}
	return sns.NewFromConfig(cfg)
}
