package accounts

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

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
