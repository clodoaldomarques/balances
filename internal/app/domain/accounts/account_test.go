package accounts

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestAccount_ChangeStatus(t *testing.T) {
	tests := []struct {
		name  string
		setup func() Account
		args  func() Status
		want  func(t *testing.T, a Account)
	}{
		{
			name: "when change status for Only Credit",
			setup: func() Account {
				return buildAccount()
			},
			args: func() Status {
				return OnlyCredit
			},
			want: func(t *testing.T, a Account) {
				assert.Equal(t, "only_credit", string(a.Status))
				assert.Equal(t, int64(2), a.Version)
			},
		},
		{
			name: "when change status for Only Debit",
			setup: func() Account {
				return buildAccount()
			},
			args: func() Status {
				return OnlyDebit
			},
			want: func(t *testing.T, a Account) {
				assert.Equal(t, "only_debit", string(a.Status))
				assert.Equal(t, int64(2), a.Version)
			},
		},
		{
			name: "when change status for Inative",
			setup: func() Account {
				return buildAccount()
			},
			args: func() Status {
				return Inative
			},
			want: func(t *testing.T, a Account) {
				assert.Equal(t, "inactive", string(a.Status))
				assert.Equal(t, int64(2), a.Version)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.setup()
			a.ChangeStatus(tt.args())
			tt.want(t, a)
		})
	}
}

func TestAccount_ChangeLimit(t *testing.T) {
	tests := []struct {
		name  string
		setup func() Account
		args  func() (string, decimal.Decimal)
		want  func(t *testing.T, a Account, e error)
	}{
		{
			name: "when change max limits with sucess",
			setup: func() Account {
				return buildAccount()
			},
			args: func() (string, decimal.Decimal) {
				return MaxLimit, decimal.NewFromInt(150)
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Nil(t, e)
				assert.Equal(t, decimal.NewFromInt(150), a.Limits[MaxLimit])
				assert.Equal(t, int64(2), a.Version)
			},
		},
		{
			name: "when change max limits with error",
			setup: func() Account {
				return buildAccount()
			},
			args: func() (string, decimal.Decimal) {
				return MaxLimit, decimal.NewFromInt(90)
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Error(t, e)
				assert.Equal(t, "new limit can not less than total limit", e.Error())
			},
		},
		{
			name: "when change total limits with sucess",
			setup: func() Account {
				return buildAccount()
			},
			args: func() (string, decimal.Decimal) {
				return TotalLimit, decimal.NewFromInt(90)
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Nil(t, e)
				assert.Equal(t, decimal.NewFromInt(90), a.Limits[TotalLimit])
				assert.Equal(t, int64(2), a.Version)
			},
		},
		{
			name: "when change total limits with error",
			setup: func() Account {
				return buildAccount()
			},
			args: func() (string, decimal.Decimal) {
				return TotalLimit, decimal.NewFromInt(150)
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Error(t, e)
				assert.Equal(t, "new limit can not great than max limit", e.Error())
			},
		},
		{
			name: "when change any limits with disabled account",
			setup: func() Account {
				a := buildAccount()
				a.ChangeStatus(Inative)
				return a
			},
			args: func() (string, decimal.Decimal) {
				return TotalLimit, decimal.NewFromInt(50)
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Error(t, e)
				assert.Equal(t, "this Account is disabled for this operation", e.Error())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.setup()
			err := a.ChangeLimit(tt.args())
			tt.want(t, a, err)
		})
	}
}

func TestAccount_ChangeBalances(t *testing.T) {
	tests := []struct {
		name  string
		setup func() Account
		args  func() []Impact
		want  func(t *testing.T, a Account, e error)
	}{
		{
			name: "credit in available balance with sucess",
			setup: func() Account {
				return buildAccount()
			},
			args: func() []Impact {
				return []Impact{
					{
						Balance:   "available_balance",
						Operation: "CREDIT",
						Amount:    decimal.NewFromInt(100),
						Rules:     []string{},
					},
				}
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Nil(t, e)
				assert.Equal(t, decimal.NewFromInt(200), a.Balances[AvailableBalance])
				assert.Equal(t, int64(2), a.Version)
			},
		},
		{
			name: "credit in savings balance with sucess",
			setup: func() Account {
				return buildAccount()
			},
			args: func() []Impact {
				return []Impact{
					{
						Balance:   "savings_balance",
						Operation: "CREDIT",
						Amount:    decimal.NewFromInt(100),
						Rules:     []string{},
					},
				}
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Nil(t, e)
				assert.Equal(t, decimal.NewFromInt(200), a.Balances[SavingsBalance])
				assert.Equal(t, int64(2), a.Version)
			},
		},
		{
			name: "credit in blocked balance with sucess",
			setup: func() Account {
				return buildAccount()
			},
			args: func() []Impact {
				return []Impact{
					{
						Balance:   "blocked_balance",
						Operation: "CREDIT",
						Amount:    decimal.NewFromInt(50),
						Rules:     []string{},
					},
				}
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Nil(t, e)
				assert.Equal(t, decimal.NewFromInt(150), a.Balances[BlockedBalance])
				assert.Equal(t, int64(2), a.Version)
			},
		},
		{
			name: "debit in available balance with sucess",
			setup: func() Account {
				return buildAccount()
			},
			args: func() []Impact {
				return []Impact{
					{
						Balance:   "available_balance",
						Operation: "DEBIT",
						Amount:    decimal.NewFromInt(50),
						Rules:     []string{"ConsiderAvailableBalance"},
					},
				}
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Nil(t, e)
				assert.Equal(t, decimal.NewFromInt(50), a.Balances[AvailableBalance])
				assert.Equal(t, int64(2), a.Version)
			},
		},
		{
			name: "debit in available balance with sucess",
			setup: func() Account {
				return buildAccount()
			},
			args: func() []Impact {
				return []Impact{
					{
						Balance:   "savings_balance",
						Operation: "DEBIT",
						Amount:    decimal.NewFromInt(50),
						Rules:     []string{"ConsiderSavingsBalance"},
					},
				}
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Nil(t, e)
				assert.Equal(t, decimal.NewFromInt(50), a.Balances[SavingsBalance])
				assert.Equal(t, int64(2), a.Version)
			},
		},
		{
			name: "debit in blocked balance with sucess",
			setup: func() Account {
				return buildAccount()
			},
			args: func() []Impact {
				return []Impact{
					{
						Balance:   "blocked_balance",
						Operation: "DEBIT",
						Amount:    decimal.NewFromInt(50),
						Rules:     []string{"ConsiderBlockedBalance"},
					},
				}
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Nil(t, e)
				assert.Equal(t, decimal.NewFromInt(50), a.Balances[BlockedBalance])
				assert.Equal(t, int64(2), a.Version)
			},
		},
		{
			name: "transfer value from available balance to blocked balance with sucess",
			setup: func() Account {
				return buildAccount()
			},
			args: func() []Impact {
				return []Impact{
					{
						Balance:   "available_balance",
						Operation: "DEBIT",
						Amount:    decimal.NewFromInt(50),
						Rules:     []string{"ConsiderAvailableBalance"},
					},
					{
						Balance:   "blocked_balance",
						Operation: "CREDIT",
						Amount:    decimal.NewFromInt(50),
						Rules:     []string{},
					},
				}
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Nil(t, e)
				assert.Equal(t, decimal.NewFromInt(50), a.Balances[AvailableBalance])
				assert.Equal(t, decimal.NewFromInt(150), a.Balances[BlockedBalance])
				assert.Equal(t, int64(2), a.Version)
			},
		},
		{
			name: "debit in available balance with insuficient balance",
			setup: func() Account {
				return buildAccount()
			},
			args: func() []Impact {
				return []Impact{
					{
						Balance:   "available_balance",
						Operation: "DEBIT",
						Amount:    decimal.NewFromInt(200),
						Rules:     []string{"ConsiderAvailableBalance"},
					},
				}
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Error(t, e)
				assert.Equal(t, "insuficient balance", e.Error())

			},
		},
		{
			name: "debit in savings balance with insuficient balance",
			setup: func() Account {
				return buildAccount()
			},
			args: func() []Impact {
				return []Impact{
					{
						Balance:   "savings_balance",
						Operation: "DEBIT",
						Amount:    decimal.NewFromInt(200),
						Rules:     []string{"ConsiderSavingsBalance"},
					},
				}
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Error(t, e)
				assert.Equal(t, "insuficient balance", e.Error())

			},
		},
		{
			name: "debit in blocked balance with insuficient balance",
			setup: func() Account {
				return buildAccount()
			},
			args: func() []Impact {
				return []Impact{
					{
						Balance:   "blocked_balance",
						Operation: "DEBIT",
						Amount:    decimal.NewFromInt(200),
						Rules:     []string{"ConsiderBlockedBalance"},
					},
				}
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Error(t, e)
				assert.Equal(t, "insuficient balance", e.Error())

			},
		},
		{
			name: "error when account is only credit",
			setup: func() Account {
				a := buildAccount()
				a.ChangeStatus(OnlyCredit)
				return a
			},
			args: func() []Impact {
				return []Impact{
					{
						Balance:   "blocked_balance",
						Operation: "DEBIT",
						Amount:    decimal.NewFromInt(50),
						Rules:     []string{"ConsiderBlockedBalance"},
					},
				}
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Error(t, e)
				assert.Equal(t, "operation invalid", e.Error())

			},
		},
		{
			name: "error when account is only debit",
			setup: func() Account {
				a := buildAccount()
				a.ChangeStatus(OnlyDebit)
				return a
			},
			args: func() []Impact {
				return []Impact{
					{
						Balance:   "blocked_balance",
						Operation: "CREDIT",
						Amount:    decimal.NewFromInt(50),
						Rules:     []string{},
					},
				}
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Error(t, e)
				assert.Equal(t, "operation invalid", e.Error())

			},
		},
		{
			name: "error when account is inative",
			setup: func() Account {
				a := buildAccount()
				a.ChangeStatus(Inative)
				return a
			},
			args: func() []Impact {
				return []Impact{
					{
						Balance:   "blocked_balance",
						Operation: "CREDIT",
						Amount:    decimal.NewFromInt(50),
						Rules:     []string{},
					},
				}
			},
			want: func(t *testing.T, a Account, e error) {
				assert.Error(t, e)
				assert.Equal(t, "operation invalid", e.Error())

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.setup()
			err := a.ChangeBalances(tt.args())
			tt.want(t, a, err)
		})
	}
}

func buildAccount() Account {
	return Account{
		AccountID: int64(23052013),
		OrgID:     "TN-12345678",
		Limits: map[string]decimal.Decimal{
			MaxLimit:       decimal.NewFromInt(100),
			TotalLimit:     decimal.NewFromInt(100),
			OverdraftLimit: decimal.NewFromInt(50),
		},
		Balances: map[string]decimal.Decimal{
			AvailableBalance: decimal.NewFromInt(100),
			SavingsBalance:   decimal.NewFromInt(100),
			BlockedBalance:   decimal.NewFromInt(100),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    Active,
		Version:   1,
	}
}
