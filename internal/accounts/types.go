package accounts

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

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

type Status string

// account status
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

func (a *Account) ChangeStatus(status Status) {
	a.Status = status
	a.IncreaseVersion()
}

func (a *Account) ChangeLimit(limit string, value decimal.Decimal) error {
	if a.Status == Inative {
		return ErrAccountDisabled
	}
	if err := limits[limit](a, value); err != nil {
		return err
	}
	a.Limits[limit] = value
	a.IncreaseVersion()
	return nil
}

func (a *Account) ChangeBalances(impacts []Impact) error {
	for _, i := range impacts {
		for _, r := range i.Rules {
			if err := rules[r](a, i.Amount); err != nil {
				return fmt.Errorf("Validation: %v", err)
			}
		}
		if err := Operations[i.Operation](a, i.Balance, i.Amount); err != nil {
			return fmt.Errorf("Validation: %v", err)
		}
	}
	a.IncreaseVersion()
	return nil
}

func (a *Account) IncreaseVersion() {
	a.Version++
	a.UpdatedAt = time.Now()
}

type Impact struct {
	Balance   string
	Operation string
	Amount    decimal.Decimal
	Rules     []string
}

var rules = map[string]func(*Account, decimal.Decimal) error{
	"ConsiderAvailableBalance": validateDebitAvailableBalance,
	"ConsiderSavingsBalance":   validateDebitSavingsBalance,
	"ConsiderBlockedBalance":   validateDebitBlockedBalance,
}

func validateDebitAvailableBalance(a *Account, amount decimal.Decimal) error {
	if a.Balances[AvailableBalance].LessThan(amount) {
		return ErrInsuficientBalance
	}
	return nil
}

func validateDebitSavingsBalance(a *Account, amount decimal.Decimal) error {
	if a.Balances[SavingsBalance].LessThan(amount) {
		return ErrInsuficientBalance
	}
	return nil
}

func validateDebitBlockedBalance(a *Account, amount decimal.Decimal) error {
	if a.Balances[BlockedBalance].LessThan(amount) {
		return ErrInsuficientBalance
	}
	return nil
}

var Operations = map[string]func(a *Account, balance string, amount decimal.Decimal) error{
	"DEBIT":  debit,
	"CREDIT": credit,
}

func credit(a *Account, balance string, amount decimal.Decimal) error {
	if a.Status != Active && a.Status != OnlyCredit {
		return errors.New("operation invalid")
	}
	a.Balances[balance] = a.Balances[balance].Add(amount)
	return nil
}

func debit(a *Account, balance string, amount decimal.Decimal) error {
	if a.Status != Active && a.Status != OnlyDebit {
		return errors.New("operation invalid")
	}
	a.Balances[balance] = a.Balances[balance].Sub(amount)
	return nil
}

var limits = map[string]func(a *Account, newValue decimal.Decimal) error{
	"max_limit":       validateChangeMaxLimit,
	"total_limit":     validateChangeTotalLimit,
	"overdraft_limit": validateChangeOverdraftLimit,
}

func validateChangeMaxLimit(a *Account, newValue decimal.Decimal) error {
	if a.Limits[TotalLimit].GreaterThan(newValue) {
		return errors.New("new limit can not less than total limit")
	}
	return nil
}

func validateChangeTotalLimit(a *Account, newValue decimal.Decimal) error {
	if a.Limits[MaxLimit].LessThan(newValue) {
		return errors.New("new limit can not great than max limit")
	}
	return nil
}

func validateChangeOverdraftLimit(a *Account, newValue decimal.Decimal) error {
	return nil
}
