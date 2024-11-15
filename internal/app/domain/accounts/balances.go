package accounts

import "github.com/shopspring/decimal"

const (
	AvailableBalance = "available_balance"
	SavingsBalance   = "savings_balance"
	BlockedBalance   = "blocked_balance"
)

const (
	ConsiderAvailableBalance = "consider_available_balance"
	ConsiderSavingsBalance   = "consider_savings_balance"
	ConsiderBlockedBalance   = "consider_blocked_balance"
)

var rules = map[string]func(*Account, decimal.Decimal) error{
	ConsiderAvailableBalance: validateDebitAvailableBalance,
	ConsiderSavingsBalance:   validateDebitSavingsBalance,
	ConsiderBlockedBalance:   validateDebitBlockedBalance,
}

func validateDebitAvailableBalance(a *Account, amount decimal.Decimal) error {
	if a.Balances[AvailableBalance].LessThan(amount) {
		return ErrInsuficientBalance{}
	}
	return nil
}

func validateDebitSavingsBalance(a *Account, amount decimal.Decimal) error {
	if a.Balances[SavingsBalance].LessThan(amount) {
		return ErrInsuficientBalance{}
	}
	return nil
}

func validateDebitBlockedBalance(a *Account, amount decimal.Decimal) error {
	if a.Balances[BlockedBalance].LessThan(amount) {
		return ErrInsuficientBalance{}
	}
	return nil
}
