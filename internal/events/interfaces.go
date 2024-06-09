package events

import (
	"context"
	"time"
)

type Repository interface {
	CreateAccountEvent(ctx context.Context, event CreateAccountEvent) error
	ChangeLimitsEvent(ctx context.Context, event ChangeLimitsEvent) error
	ChangeBalanceEvent(ctx context.Context, event ChangeBalanceEvent) error
	ChangeStatusEvent(ctx context.Context, event ChangeStatusEvent) error
}

type Event interface {
	GetTrackingID() string
	GetEventType() string
	GetAccountID() int64
	GetTenantID() string
	GetEventDate() time.Time
	GetVersion() int64
	GetEventData() string
}
