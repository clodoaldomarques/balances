package mysqldb

import (
	"balances/internal/app/domain/accounts"
	"balances/internal/app/domain/commons"
	"encoding/json"
	"time"
)

type AccountTable struct {
	AccountID int64              `json:"account_id,omitempty"`
	OrgID     string             `json:"org_id,omitempty"`
	Limits    commons.DecimalMap `json:"limits,omitempty"`
	Balances  commons.DecimalMap `json:"balances,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
	Status    string             `json:"status,omitempty"`
	Version   int64              `json:"version,omitempty"`
}

func (a AccountTable) ToEntity() accounts.Account {
	return accounts.Account{
		AccountID: a.AccountID,
		OrgID:     a.OrgID,
		Limits:    a.Limits,
		Balances:  a.Balances,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		Status:    accounts.Status(a.Status),
		Version:   a.Version,
	}
}

func AccountToTable(a accounts.Account) AccountTable {
	return AccountTable{
		AccountID: a.AccountID,
		OrgID:     a.OrgID,
		Limits:    a.Limits,
		Balances:  a.Balances,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		Status:    string(a.Status),
		Version:   a.Version,
	}
}

type EntryTable struct {
	TrackingID string    `json:"tracking_id" validate:"required"`
	AccountID  int64     `json:"account_id" validate:"required"`
	OrgID      string    `json:"org_id" validate:"required"`
	Impacts    []byte    `json:"impacts" validate:"required"`
	CreatedAt  time.Time `json:"created_at" validate:"required"`
}

func EntryToTable(e accounts.Entry) (EntryTable, error) {
	impacts, err := json.Marshal(e.Impacts)
	if err != nil {
		return EntryTable{}, err
	}
	return EntryTable{
		TrackingID: e.TrackingID,
		AccountID:  e.AccountID,
		OrgID:      e.OrgID,
		Impacts:    impacts,
		CreatedAt:  e.CreatedAt,
	}, nil
}
