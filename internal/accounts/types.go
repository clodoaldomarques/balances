package accounts

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Status string

const (
	Active     Status = "active"
	OnlyCredit Status = "only_credit"
	OnlyDebit  Status = "only_debit"
	Inative    Status = "inactive"
)

// accounts limits
const (
	MaxLimit       = "max_limit"
	TotalLimit     = "total_limit"
	OverdraftLimit = "overdraft_limit"
)

// availables balances
const (
	AvailableBalance = "available_balance"
	SavingsBalance   = "savings_balance"
	BlockedBalance   = "blocked_balance"
)

type Account struct {
	AccountID int64      `json:"account_id,omitempty"`
	TenantID  string     `json:"tenant_id,omitempty"`
	Limits    DecimalMap `json:"limits,omitempty"`
	Balances  DecimalMap `json:"balances,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	Status    Status     `json:"status,omitempty"`
	Version   int64      `json:"version,omitempty"`
}

func init() {
	decimal.MarshalJSONWithoutQuotes = true
}

type DecimalMap map[string]decimal.Decimal

func (d DecimalMap) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *DecimalMap) Scan(src interface{}) error {
	switch data := src.(type) {
	case []byte:
		return json.Unmarshal(data, d)
	case string:
		return json.Unmarshal([]byte(data), d)
	default:
		return fmt.Errorf("unsupported type: %T", src)
	}
}
