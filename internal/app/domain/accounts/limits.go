package accounts

import "github.com/shopspring/decimal"

const (
	MaxLimit       = "max_limit"
	TotalLimit     = "total_limit"
	OverdraftLimit = "overdraft_limit"
)

var limits = map[string]func(a *Account, newValue decimal.Decimal) error{
	MaxLimit:       validateChangeMaxLimit,
	TotalLimit:     validateChangeTotalLimit,
	OverdraftLimit: validateChangeOverdraftLimit,
}

func validateChangeMaxLimit(a *Account, newValue decimal.Decimal) error {
	if a.Limits[TotalLimit].GreaterThan(newValue) {
		return ErrValidateLimit{msg: "new limit can not less than total limit"}
	}
	return nil
}

func validateChangeTotalLimit(a *Account, newValue decimal.Decimal) error {
	if a.Limits[MaxLimit].LessThan(newValue) {
		return ErrValidateLimit{msg: "new limit can not great than max limit"}
	}
	return nil
}

func validateChangeOverdraftLimit(a *Account, newValue decimal.Decimal) error {
	return nil
}
