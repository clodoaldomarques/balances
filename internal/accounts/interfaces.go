package accounts

import (
	"balances/internal/events"
	"context"
)

type AccountRepo interface {
	CreateNewAccount(ctx context.Context, account Account) error
	RetrieveAccountByID(ctx context.Context, accountID int64, orgID string) (Account, error)
	UpdateAccount(ctx context.Context, account Account) error
}

type Publisher interface {
	Notify(event events.Event)
}
