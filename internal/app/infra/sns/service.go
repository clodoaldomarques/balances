package sns

import (
	"context"

	"github.com/clodoaldomarques/balances/configs"
	"github.com/clodoaldomarques/balances/internal/shared/domain/events"
	"github.com/clodoaldomarques/balances/pkg/logger"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type Publisher struct {
	ctx      context.Context
	svc      *sns.Client
	topicARN string
}

func NewPublisher(ctx context.Context) *Publisher {
	return &Publisher{
		ctx:      ctx,
		svc:      NewSNSClient(ctx),
		topicARN: configs.New().BalancesSNSTopic,
	}
}

func (p Publisher) Emit(ctx context.Context, evt events.Event) {
	input := &sns.PublishInput{
		Message:  evt.ToMessage(),
		TopicArn: &p.topicARN,
	}

	result, err := p.svc.Publish(ctx, input)
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
