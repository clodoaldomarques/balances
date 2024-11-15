package mysqldb

import (
	"balances/internal/app/domain/accounts"
	"balances/internal/app/domain/commons"
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
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

type ImpactTable struct {
	Balance   string          `json:"balance" validate:"required"`
	Operation string          `json:"operation" validate:"required"`
	Amount    decimal.Decimal `json:"amount" validate:"required"`
	Rules     []string        `json:"rules,omitempty"`
}

func impactToTable(impacts []accounts.Impact) []ImpactTable {
	impactsTable := make([]ImpactTable, 0, len(impacts))
	for _, i := range impacts {
		it := ImpactTable{
			Balance:   i.Balance,
			Operation: i.Operation,
			Amount:    i.Amount,
			Rules:     i.Rules,
		}
		impactsTable = append(impactsTable, it)
	}
	return impactsTable
}

func toByte(impactsTable []ImpactTable) ([]byte, error) {
	impacts, err := json.Marshal(impactsTable)
	if err != nil {
		return nil, err
	}
	return impacts, nil
}

func EntryToTable(e accounts.Entry) (EntryTable, error) {
	impacts, err := toByte(impactToTable(e.Impacts))
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
