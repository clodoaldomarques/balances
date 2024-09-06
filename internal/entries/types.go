package entries

import "github.com/shopspring/decimal"

type Entry struct {
	AccountID int64    `json:"account_id" validate:"required"`
	TenantID  string   `json:"tenant_id" validate:"required"`
	Impacts   []Impact `json:"impacts" validate:"required"`
}

type Impact struct {
	Balance   string          `json:"balance" validate:"required"`
	Operation string          `json:"operation" validate:"required"`
	Amount    decimal.Decimal `json:"amount" validate:"required"`
	Rules     []string        `json:"rules,omitempty"`
}
