package accounts

import (
	"context"

	"github.com/shopspring/decimal"
)

type Service struct {
	rep Repository
	pub Topic
}

func NewService(r Repository, p Topic) *Service {
	return &Service{
		rep: r,
		pub: p,
	}
}

func (s Service) CreateNewAccount(ctx context.Context, a Account) (Account, error) {
	if err := s.rep.SaveNewAccount(ctx, a); err != nil {
		return Account{}, err
	}

	evt := buildCreateAccountEvent(a)
	s.pub.Emit(ctx, evt)

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

	evt := buildUpdateAccountEvent(acc)
	s.pub.Emit(ctx, evt)

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

	evt := buildUpdateAccountEvent(acc)
	s.pub.Emit(ctx, evt)

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

	evt := buildProcessEntryEvent(acc, e)
	s.pub.Emit(ctx, evt)

	return acc, nil
}

func buildCreateAccountEvent(a Account) Event {
	return CreateAccountEvent{
		AccountID: a.AccountID,
		OrgID:     a.OrgID,
		Limits:    a.Limits,
		Balances:  a.Balances,
		CreatedAt: a.CreatedAt,
		Status:    string(a.Status),
		Version:   a.Version,
	}
}

func buildUpdateAccountEvent(a Account) Event {
	return UpdateAccountEvent{
		AccountID: a.AccountID,
		OrgID:     a.OrgID,
		Limits:    a.Limits,
		Balances:  a.Balances,
		UpdatedAt: a.UpdatedAt,
		Status:    string(a.Status),
		Version:   a.Version,
	}
}

func buildProcessEntryEvent(a Account, e Entry) Event {
	return ProcessEntryEvent{
		AccountID:  a.AccountID,
		OrgID:      a.OrgID,
		TrackingID: e.TrackingID,
		Impacts:    buildImpactEvents(e.Impacts),
		Limits:     a.Limits,
		Balances:   a.Balances,
		Version:    a.Version,
		CreatedAt:  e.CreatedAt,
	}
}

func buildImpactEvents(impacts []Impact) []ImpactEvent {
	evts := make([]ImpactEvent, 0, len(impacts))
	for _, i := range impacts {
		new := ImpactEvent{
			Balance:   i.Balance,
			Operation: i.Operation,
			Amount:    i.Amount,
			Rules:     i.Rules,
		}
		evts = append(evts, new)
	}
	return evts
}
