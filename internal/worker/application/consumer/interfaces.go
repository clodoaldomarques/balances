package consumer

import (
	"balances/internal/app/domain/events"
	"context"
)

type Queue interface {
	Retrieve(ctx context.Context) (map[string]events.Event, error)
	DeleteMessages(receiptHandles []string) error
}
