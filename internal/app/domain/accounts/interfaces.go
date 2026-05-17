package accounts

import (
	"balances/internal/shared/domain/events"
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=mock.go -package=accounts
type Repository interface {
	SaveNewAccount(ctx context.Context, a Account) error
	UpdateExistingAccount(ctx context.Context, a Account) error
	RetrieveAccountByID(ctx context.Context, accountID int64, orgID string) (Account, error)
	SaveEntryAndUpdateAccount(ctx context.Context, e Entry, a Account) error
}

type Publisher interface {
	Emit(ctx context.Context, e events.Event)
}
