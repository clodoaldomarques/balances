package consumer

import (
	"balances/internal/shared/domain/events"
	"balances/internal/worker/domain/daily"
	"balances/internal/worker/infra/database/mysqldb"
	"balances/internal/worker/infra/sqs"
	"balances/pkg/logger"
	"context"
)

type Consumer struct {
	sqs  Queue
	serv daily.Service
}

func New(ctx context.Context) *Consumer {
	return &Consumer{
		sqs:  sqs.New(ctx, 10, 1),
		serv: *daily.NewService(mysqldb.NewRepository(ctx)),
	}
}

func (c Consumer) Start() {
	ctx := context.Background()
	for {
		evts, err := c.sqs.Retrieve(ctx)
		if err != nil {
			logger.Error(ctx, err.Error(), logger.Fields{})
		}
		receiptHandles := make([]string, 0, len(evts))
		for k, v := range evts {
			receiptHandles = append(receiptHandles, k)
			if err := c.process(ctx, v); err != nil {
				logger.Error(ctx, err.Error(), logger.Fields{})
			}
		}

		if len(receiptHandles) > 0 {
			if err := c.sqs.DeleteMessages(receiptHandles); err != nil {
				logger.Error(ctx, err.Error(), logger.Fields{})
			}
		}

	}
}

func (c Consumer) process(ctx context.Context, e events.Event) error {
	logger.Info(ctx, "event processed", logger.Fields{
		"event_id":   e.EventID,
		"event_type": e.EventType,
		"event_data": e.Data,
		"event_date": e.EventDate,
	})
	return nil
}
