package mysqldb

import (
	"balances/internal/app/domain/accounts"
	"context"
	"time"

	"github.com/shopspring/decimal"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r Repository) SaveNewAccount(ctx context.Context, a accounts.Account) error {
	return nil
}
func (r Repository) UpdateExistingAccount(ctx context.Context, a accounts.Account) error {
	return nil
}
func (r Repository) RetrieveAccountByID(ctx context.Context, accountID int64) (accounts.Account, error) {
	return accounts.Account{
		AccountID: int64(23052013),
		OrgID:     "TN-12345678",
		Limits: map[string]decimal.Decimal{
			accounts.MaxLimit:       decimal.NewFromInt(100),
			accounts.TotalLimit:     decimal.NewFromInt(100),
			accounts.OverdraftLimit: decimal.NewFromInt(50),
		},
		Balances: map[string]decimal.Decimal{
			accounts.AvailableBalance: decimal.NewFromInt(100),
			accounts.SavingsBalance:   decimal.NewFromInt(100),
			accounts.BlockedBalance:   decimal.NewFromInt(100),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    accounts.Active,
		Version:   1,
	}, nil
}

func (r Repository) SaveEntryAndUpdateAccount(ctx context.Context, e accounts.Entry, a accounts.Account) error {
	return nil
}
