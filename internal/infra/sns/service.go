package sns

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/clodoaldomarques/balances-api/config"
	"github.com/clodoaldomarques/balances-api/internal/domain/accounts"
	"github.com/clodoaldomarques/core-sdk/pkg/logger"
	"github.com/google/uuid"
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
		topicARN: config.New().BalancesSNSTopic,
	}
}

func (p Publisher) Emit(ctx context.Context, evt accounts.Event) {
	e := Event{
		EventID:   uuid.New(),
		EventType: evt.EventType(),
		EventData: evt.EventData(),
		EventDate: time.Now(),
	}

	input := &sns.PublishInput{
		Message:  e.ToMessage(),
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
