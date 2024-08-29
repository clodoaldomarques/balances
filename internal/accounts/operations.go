package accounts

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"
)

type Service struct {
	repo AccountRepo
	publ Publisher
}

func NewService(rep AccountRepo, pub Publisher) *Service {
	return &Service{
		repo: rep,
		publ: pub,
	}
}

func (s Service) CreateAccount(ctx context.Context, acc Account) error {
	if err := s.repo.CreateNewAccount(ctx, acc); err != nil {
		return err
	}
	return nil
}

func (s Service) UpdateAccountLimits(ctx context.Context, acc Account) error {
	account, err := s.repo.RetrieveAccountByID(ctx, acc.AccountID, acc.TenantID)
	if err != nil {
		return fmt.Errorf("can't update account: %v", err)
	}

	if err := account.UpdateAccountLimits(acc.Limits); err != nil {
		return fmt.Errorf("can´t update account limits: %v", err)
	}

	return nil
}

func (s Service) UpdateAccountStatus(ctx context.Context, accountID int64, tenantID string, status string) error {
	account, err := s.repo.RetrieveAccountByID(ctx, accountID, tenantID)
	if err != nil {
		return fmt.Errorf("can't update account: %v", err)
	}

	if err := account.UpdateAccountStatus(Status(status)); err != nil {
		return fmt.Errorf("can't update account: %v", err)
	}

	return nil
}

func (s Service) UpdateAccountBalances(ctx context.Context, accountID int64, tenantID string, operation string, balances map[string]decimal.Decimal, amount decimal.Decimal) error {
	account, err := s.repo.RetrieveAccountByID(ctx, accountID, tenantID)
	if err != nil {
		return fmt.Errorf("can't update account: %v", err)
	}
	for b, _ := range balances {
		account.UpdateAccountBalances(operation, b, amount, nil)
	}

	return nil
}

func (s Service) RetrieveAccountByID(ctx context.Context, accountID int64, tenantID string) (Account, error) {
	return s.repo.RetrieveAccountByID(ctx, accountID, tenantID)
}
