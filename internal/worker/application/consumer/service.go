package consumer

import (
	"balances/internal/worker/infra/sqs"
	"context"
)

type Consumer struct {
	sqs *sqs.Handler
}

func New(ctx context.Context) *Consumer {
	return &Consumer{
		sqs: sqs.New(ctx, 10, 1),
	}
}

func Start() {

}
