package sns

import (
	"balances/configs"
	"balances/internal/app/domain/events"
	"balances/pkg/logger"
	"context"

	"github.com/aws/aws-sdk-go/service/sns"
)

type Publisher struct {
	ctx      context.Context
	svc      *sns.SNS
	topicARN string
}

func NewPublisher(ctx context.Context) *Publisher {
	svc := NewSNSClient(ctx)
	return &Publisher{
		ctx:      ctx,
		svc:      svc,
		topicARN: configs.New().BalancesSNSTopic,
	}
}

func (p Publisher) Emit(ctx context.Context, evt events.Event) {
	input := &sns.PublishInput{
		Message:  evt.ToMessage(),
		TopicArn: &p.topicARN,
	}

	result, err := p.svc.PublishWithContext(ctx, input)
	if err != nil {
		logger.Error(ctx, "error on publish", logger.Fields{
			"error": err.Error(),
		})
	}
	logger.Info(ctx, "event published with success", logger.Fields{
		"message_id": result.MessageId,
		"event":      evt,
	})
}
