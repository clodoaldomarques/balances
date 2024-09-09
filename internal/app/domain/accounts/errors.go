package accounts

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type ErrInsuficientBalance struct {
	Msg string
}

func (e ErrInsuficientBalance) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return "insuficient balance"
}

type ErrAccountDisabled struct {
	Msg string
}

func (e ErrAccountDisabled) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return "this Account is disabled for this operation"
}

type ErrValidateLimit struct {
	Msg string
}

func (e ErrValidateLimit) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return "this is an invalid operation"
}

type ErrValidateOperation struct {
	Msg       string
	Operation string
	Balance   string
	Amount    decimal.Decimal
}

func (e ErrValidateOperation) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return fmt.Sprintf("invalid operation: %s, balance: %s, amount: %d", e.Operation, e.Balance, e.Amount)
}
