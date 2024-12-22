package daily

import (
	"context"
	"time"
)

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{
		r: r,
	}
}

func (s Service) CreateNewBalance(ctx context.Context, b Balance) error {
	return s.r.SaveNewBalance(ctx, b)
}

func (s Service) UpdateExistingBalance(ctx context.Context, b Balance) error {
	old, err := s.r.RetrieveLastBalance(ctx, b.AccountID, b.OrgID)
	if err != nil {
		return err
	}

	if old.Date != b.Date {
		return s.FillMissingBalances(ctx, old, b)
	}

	return s.r.UpdateExistingBalance(ctx, b)
}

func (s Service) FillMissingBalances(ctx context.Context, old, new Balance) error {

	return nil
}

func (s Service) RetrieveBalanceByDate(ctx context.Context, accountID int64, orgID string, initialDate, finalDate time.Time) ([]Balance, error) {
	return nil, nil
}
