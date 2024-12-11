package sns

import (
	"balances/pkg/aws"
	"balances/pkg/logger"
	"context"

	"github.com/aws/aws-sdk-go/service/sns"
)

func NewSNSClient(ctx context.Context) *sns.SNS {
	sess, err := aws.NewAWSConfig(ctx)
	if err != nil {
		logger.Fatal(ctx, err.Error(), logger.Fields{})
	}
	return sns.New(sess)
}
