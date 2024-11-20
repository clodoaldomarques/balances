package accounts

import (
	"context"

	"github.com/shopspring/decimal"
)

type Service struct {
	rep Repository
}

func NewService(r Repository) *Service {
	return &Service{
		rep: r,
	}
}

func (s Service) CreateNewAccount(ctx context.Context, a Account) (Account, error) {
	if err := s.rep.SaveNewAccount(ctx, a); err != nil {
		return Account{}, err
	}
	return a, nil
}

func (s Service) UpdateAccountLimits(ctx context.Context, accountID int64, orgID string, limits map[string]decimal.Decimal) (Account, error) {
	acc, err := s.rep.RetrieveAccountByID(ctx, accountID, orgID)
	if err != nil {
		return Account{}, err
	}

	for limit, value := range limits {
		if err = acc.ChangeLimit(limit, value); err != nil {
			return Account{}, err
		}
	}

	if err = s.rep.UpdateExistingAccount(ctx, acc); err != nil {
		return Account{}, err
	}

	return acc, nil

}

func (s Service) UpdateAccountStatus(ctx context.Context, accountID int64, orgID string, status Status) (Account, error) {
	acc, err := s.rep.RetrieveAccountByID(ctx, accountID, orgID)
	if err != nil {
		return Account{}, err
	}

	acc.ChangeStatus(status)

	if err = s.rep.UpdateExistingAccount(ctx, acc); err != nil {
		return Account{}, err
	}

	return acc, nil
}

func (s Service) ProcessEntry(ctx context.Context, e Entry) (Account, error) {
	acc, err := s.rep.RetrieveAccountByID(ctx, e.AccountID, e.OrgID)
	if err != nil {
		return Account{}, err
	}

	if err = acc.ChangeBalances(e.Impacts); err != nil {
		return Account{}, err
	}

	if err = s.rep.SaveEntryAndUpdateAccount(ctx, e, acc); err != nil {
		return Account{}, err
	}

	return acc, nil
}
