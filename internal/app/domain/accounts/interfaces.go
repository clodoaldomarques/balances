package accounts

import "context"

//go:mockgen -source=interfaces.go -destination=mock.go -package=accounts
type Repository interface {
	SaveNewAccount(ctx context.Context, a Account) error
	UpdateExistingAccount(ctx context.Context, a Account) error
	RetrieveAccountByID(ctx context.Context, accountID int64) (Account, error)
	SaveEntryAndUpdateAccount(ctx context.Context, e Entry, a Account) error
}
