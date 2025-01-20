package consumer

import (
	"balances/internal/shared/domain/events"
	"balances/internal/worker/domain/daily"
	"balances/internal/worker/infra/database/mysqldb"
	"balances/internal/worker/infra/sqs"
	"balances/pkg/logger"
	"context"
	"encoding/json"
	"fmt"
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

var eventFunc = map[events.Type]func(context.Context, daily.Service, events.Event) error{
	daily.CREATE_ACCOUNT: processCreateAccount,
	daily.UPDATE_ACCOUNT: processUpdateAccount,
	daily.PROCESS_ENTRY:  processProcessEntry,
}

func processCreateAccount(ctx context.Context, s daily.Service, e events.Event) error {
	var cae daily.CreateAccountEvent
	if err := json.Unmarshal([]byte(fmt.Sprint(e.Data)), &cae); err != nil {
		return err
	}

	if err := s.CreateNewBalance(ctx, cae.ToDailyBalance()); err != nil {
		return err
	}
	return nil
}

func processUpdateAccount(ctx context.Context, s daily.Service, e events.Event) error {
	var uae daily.UpdateAccountEvent
	if err := json.Unmarshal([]byte(fmt.Sprint(e.Data)), &uae); err != nil {
		return err
	}

	if err := s.UpdateExistingBalance(ctx, uae.ToDailyBalance()); err != nil {
		return err
	}
	return nil
}

func processProcessEntry(ctx context.Context, s daily.Service, e events.Event) error {
	var pee daily.ProcessEntryEvent
	dt := fmt.Sprint(e.Data)
	if err := json.Unmarshal([]byte(dt), &pee); err != nil {
		return err
	}

	if err := s.UpdateExistingBalance(ctx, pee.ToDailyBalance()); err != nil {
		return err
	}
	return nil
}

func (c Consumer) process(ctx context.Context, e events.Event) error {
	logger.Info(ctx, "event processed", logger.Fields{
		"event_id":   e.EventID,
		"event_type": e.EventType,
		"event_data": e.Data,
		"event_date": e.EventDate,
	})

	if err := eventFunc[e.EventType](ctx, c.serv, e); err != nil {
		return err
	}

	return nil
}
