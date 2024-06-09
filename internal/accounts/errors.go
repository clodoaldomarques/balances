package accounts

import (
	"errors"
)

var (
	ErrInsuficientBalance = errors.New("insuficient balance")
	ErrAccountDisabled    = errors.New("this Account is disabled for this operation")
)
