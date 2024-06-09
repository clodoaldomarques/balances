package accounts

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestAccount_UpdateAccountLimits(t *testing.T) {
	type fields struct {
		AccountID int64
		TenantID  string
		Limits    DecimalMap
		Balances  DecimalMap
		CreatedAt time.Time
		UpdatedAt time.Time
		Status    Status
		Version   int64
	}
	type args struct {
		limits map[string]decimal.Decimal
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				AccountID: tt.fields.AccountID,
				TenantID:  tt.fields.TenantID,
				Limits:    tt.fields.Limits,
				Balances:  tt.fields.Balances,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				Status:    tt.fields.Status,
				Version:   tt.fields.Version,
			}
			if err := a.UpdateAccountLimits(tt.args.limits); (err != nil) != tt.wantErr {
				t.Errorf("Account.UpdateAccountLimits() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccount_UpdateAccountStatus(t *testing.T) {
	type fields struct {
		AccountID int64
		TenantID  string
		Limits    DecimalMap
		Balances  DecimalMap
		CreatedAt time.Time
		UpdatedAt time.Time
		Status    Status
		Version   int64
	}
	type args struct {
		status Status
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				AccountID: tt.fields.AccountID,
				TenantID:  tt.fields.TenantID,
				Limits:    tt.fields.Limits,
				Balances:  tt.fields.Balances,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				Status:    tt.fields.Status,
				Version:   tt.fields.Version,
			}
			if err := a.UpdateAccountStatus(tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("Account.UpdateAccountStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccount_UpdateAccountBalances(t *testing.T) {
	type fields struct {
		AccountID int64
		TenantID  string
		Limits    DecimalMap
		Balances  DecimalMap
		CreatedAt time.Time
		UpdatedAt time.Time
		Status    Status
		Version   int64
	}
	type args struct {
		op      string
		balance string
		amount  decimal.Decimal
		rules   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				AccountID: tt.fields.AccountID,
				TenantID:  tt.fields.TenantID,
				Limits:    tt.fields.Limits,
				Balances:  tt.fields.Balances,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				Status:    tt.fields.Status,
				Version:   tt.fields.Version,
			}
			if err := a.UpdateAccountBalances(tt.args.op, tt.args.balance, tt.args.amount, tt.args.rules); (err != nil) != tt.wantErr {
				t.Errorf("Account.UpdateAccountBalances() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
