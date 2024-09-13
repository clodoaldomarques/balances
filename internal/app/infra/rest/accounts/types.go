package accounts

import (
	"balances/internal/app/domain/accounts"
	"balances/internal/app/domain/commons"
	"time"
)

type PostAccountRequest struct {
	AccountID int64              `json:"account_id,omitempty"`
	OrgID     string             `json:"org_id,omitempty"`
	Limits    commons.DecimalMap `json:"limits,omitempty"`
	Balances  commons.DecimalMap `json:"balances,omitempty"`
}

func (p PostAccountRequest) ToEntity() accounts.Account {
	return accounts.Account{
		AccountID: p.AccountID,
		OrgID:     p.OrgID,
		Limits:    p.Limits,
		Balances:  p.Balances,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    accounts.Active,
		Version:   1,
	}
}

type PostAccountResponse struct {
	AccountID int64              `json:"account_id,omitempty"`
	OrgID     string             `json:"org_id,omitempty"`
	Limits    commons.DecimalMap `json:"limits,omitempty"`
	Balances  commons.DecimalMap `json:"balances,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
	Status    string             `json:"status,omitempty"`
	Version   int64              `json:"version,omitempty"`
}

func AccountToPostAccountResponse(acc accounts.Account) PostAccountResponse {
	return PostAccountResponse{
		AccountID: acc.AccountID,
		OrgID:     acc.OrgID,
		Limits:    acc.Limits,
		Balances:  acc.Balances,
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
		Status:    string(acc.Status),
		Version:   acc.Version,
	}
}

type PutAccountRequest struct {
	Limits commons.DecimalMap `json:"limits,omitempty"`
	Status string             `json:"status,omitempty"`
}

type PutAccountResponse struct {
	AccountID int64              `json:"account_id,omitempty"`
	OrgID     string             `json:"org_id,omitempty"`
	Limits    commons.DecimalMap `json:"limits,omitempty"`
	Balances  commons.DecimalMap `json:"balances,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
	Status    string             `json:"status,omitempty"`
	Version   int64              `json:"version,omitempty"`
}

func AccountToPutAccountResponse(acc accounts.Account) PutAccountResponse {
	return PutAccountResponse{
		AccountID: acc.AccountID,
		OrgID:     acc.OrgID,
		Limits:    acc.Limits,
		Balances:  acc.Balances,
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
		Status:    string(acc.Status),
		Version:   acc.Version,
	}
}

type PostEntryRequest struct {
	TrackingID string            `json:"tracking_id" validate:"required"`
	AccountID  int64             `json:"account_id" validate:"required"`
	OrgID      string            `json:"org_id" validate:"required"`
	Impacts    []accounts.Impact `json:"impacts" validate:"required"`
}

func (p PostEntryRequest) ToEntity() accounts.Entry {
	return accounts.Entry{
		TrackingID: p.TrackingID,
		AccountID:  p.AccountID,
		OrgID:      p.OrgID,
		Impacts:    p.Impacts,
		CreatedAt:  time.Now(),
	}
}
