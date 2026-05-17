package consumer

import (
	"balances/internal/shared/domain/events"
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=mock.go -package=consumer
type Queue interface {
	Retrieve(ctx context.Context) (map[string]events.Event, error)
	DeleteMessages(receiptHandles []string) error
}
