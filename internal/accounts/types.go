package accounts

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
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
	OverdraftLimit = "overdraft_limit"
)

var limits = map[string]string{
	MaxLimit:       MaxLimit,
	OverdraftLimit: OverdraftLimit,
}

// availables balances
const (
	AvailableBalance = "available_balance"
	SavingsBalance   = "savings_balance"
	BlockedBalance   = "blocked_balance"
)

var balances = map[string]string{
	AvailableBalance: AvailableBalance,
	SavingsBalance:   SavingsBalance,
	BlockedBalance:   BlockedBalance,
}

type rule = func(*Account, decimal.Decimal) error

var considers = map[string]rule{
	"AvailableBalance": validateDebitAvailableBalance,
	"SavingsBalance":   validateDebitSavingsBalance,
	"BlockedBalance":   validateDebitBlockedBalance,
}

func validateDebitAvailableBalance(a *Account, amount decimal.Decimal) error {
	if a.Balances[AvailableBalance].LessThan(amount) {
		return errors.New("available balance can not less than amount")
	}
	return nil
}

func validateDebitSavingsBalance(a *Account, amount decimal.Decimal) error {
	if a.Balances[SavingsBalance].LessThan(amount) {
		return errors.New("savings balance can not less than amount")
	}
	return nil
}

func validateDebitBlockedBalance(a *Account, amount decimal.Decimal) error {
	if a.Balances[BlockedBalance].LessThan(amount) {
		return errors.New("blocked balance can not less than amount")
	}
	return nil
}

type operationType = func(a *Account, balance string, amount decimal.Decimal, rules []string) error

var Operation = map[string]operationType{
	"DEBIT":  debit,
	"CREDIT": credit,
}

func credit(a *Account, balance string, amount decimal.Decimal, rules []string) error {
	if a.Status != Active && a.Status != OnlyCredit {
		return errors.New("operation invalid")
	}
	if len(rules) > 0 {
		return errors.New("is not necessary rules for credit")
	}
	a.Balances[balance] = a.Balances[balance].Add(amount)
	return nil
}

func debit(a *Account, balance string, amount decimal.Decimal, rules []string) error {
	if a.Status != Active && a.Status != OnlyDebit {
		return errors.New("operation invalid")
	}
	for _, r := range rules {
		if err := considers[r](a, amount); err != nil {
			return fmt.Errorf("Validation: %v", err)
		}
	}
	a.Balances[balance] = a.Balances[balance].Sub(amount)
	return nil
}

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

func (a *Account) UpdateAccountLimits(limits map[string]decimal.Decimal) error {
	for key, value := range limits {
		a.Limits[key] = value
	}

	a.UpdatedAt = time.Now().UTC()

	return nil
}

func (a *Account) UpdateAccountStatus(status Status) error {
	a.Status = status
	a.UpdatedAt = time.Now().UTC()
	return nil
}

func (a *Account) UpdateAccountBalances(op string, balance string, amount decimal.Decimal, rules []string) error {
	if err := Operation[op](a, balance, amount, rules); err != nil {
		return fmt.Errorf("operation invalid: %v", err)
	}
	a.UpdatedAt = time.Now().UTC()
	return nil
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
