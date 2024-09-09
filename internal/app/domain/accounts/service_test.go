package accounts

import (
	"balances/internal/app/domain/commons"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	tests := []struct {
		name  string
		setup func(ctrl *gomock.Controller) Repository
		want  func(t *testing.T, s *Service)
	}{
		{
			name: "when create new service with success",
			setup: func(ctrl *gomock.Controller) Repository {
				return NewMockRepository(ctrl)
			},
			want: func(t *testing.T, s *Service) {
				assert.NotEmpty(t, s)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			r := tt.setup(ctrl)
			tt.want(t, NewService(r))
		})
	}
}

func TestService_CreateNewAccount(t *testing.T) {
	tests := []struct {
		name  string
		setup func(ctrl *gomock.Controller) *Service
		args  func() Account
		want  func(t *testing.T, e error)
	}{
		{
			name: "when create account with sucess",
			setup: func(ctrl *gomock.Controller) *Service {
				rep := NewMockRepository(ctrl)
				rep.EXPECT().SaveNewAccount(gomock.Any(), gomock.Any()).Return(nil).Times(1)
				return NewService(rep)
			},
			args: func() Account {
				return buildAccount()
			},
			want: func(t *testing.T, e error) {
				assert.Nil(t, e)
			},
		},
		{
			name: "when repository send an error",
			setup: func(ctrl *gomock.Controller) *Service {
				rep := NewMockRepository(ctrl)
				rep.EXPECT().SaveNewAccount(gomock.Any(), gomock.Any()).Return(errors.New("repository not found")).Times(1)
				return NewService(rep)
			},
			args: func() Account {
				return buildAccount()
			},
			want: func(t *testing.T, e error) {
				assert.NotNil(t, e)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			s := tt.setup(ctrl)
			tt.want(t, s.CreateNewAccount(context.Background(), tt.args()))
		})
	}
}

func TestService_UpdateAccountLimits(t *testing.T) {
	tests := []struct {
		name  string
		setup func(ctrl *gomock.Controller) *Service
		args  func() (int64, commons.DecimalMap)
		want  func(t *testing.T, e error)
	}{
		{
			name: "when update a limit with sucess",
			setup: func(ctrl *gomock.Controller) *Service {
				rep := NewMockRepository(ctrl)
				acc := buildAccount()
				rep.EXPECT().RetrieveAccountByID(gomock.Any(), int64(230513)).Return(acc, nil).Times(1)
				rep.EXPECT().UpdateExistingAccount(gomock.Any(), gomock.Any()).Return(nil).Times(1)
				return NewService(rep)
			},
			args: func() (int64, commons.DecimalMap) {
				return int64(230513), commons.DecimalMap{
					MaxLimit: decimal.NewFromInt(150),
				}
			},
			want: func(t *testing.T, e error) {
				assert.Nil(t, e)
			},
		},
		{
			name: "when account not found",
			setup: func(ctrl *gomock.Controller) *Service {
				rep := NewMockRepository(ctrl)
				acc := Account{}
				rep.EXPECT().RetrieveAccountByID(gomock.Any(), int64(230513)).Return(acc, errors.New("Account not found")).Times(1)
				return NewService(rep)
			},
			args: func() (int64, commons.DecimalMap) {
				return int64(230513), commons.DecimalMap{
					MaxLimit: decimal.NewFromInt(150),
				}
			},
			want: func(t *testing.T, e error) {
				assert.NotNil(t, e)
				assert.Equal(t, "Account not found", e.Error())
			},
		},
		{
			name: "when send invalid limit",
			setup: func(ctrl *gomock.Controller) *Service {
				rep := NewMockRepository(ctrl)
				acc := buildAccount()
				rep.EXPECT().RetrieveAccountByID(gomock.Any(), int64(230513)).Return(acc, nil).Times(1)
				return NewService(rep)
			},
			args: func() (int64, commons.DecimalMap) {
				return int64(230513), commons.DecimalMap{
					TotalLimit: decimal.NewFromInt(150),
				}
			},
			want: func(t *testing.T, e error) {
				assert.NotNil(t, e)
				assert.Equal(t, "new limit can not great than max limit", e.Error())
			},
		},
		{
			name: "when send an error on update account",
			setup: func(ctrl *gomock.Controller) *Service {
				rep := NewMockRepository(ctrl)
				acc := buildAccount()
				rep.EXPECT().RetrieveAccountByID(gomock.Any(), int64(230513)).Return(acc, nil).Times(1)
				rep.EXPECT().UpdateExistingAccount(gomock.Any(), gomock.Any()).Return(errors.New("error on update account")).Times(1)
				return NewService(rep)
			},
			args: func() (int64, commons.DecimalMap) {
				return int64(230513), commons.DecimalMap{
					MaxLimit: decimal.NewFromInt(150),
				}
			},
			want: func(t *testing.T, e error) {
				assert.NotNil(t, e)
				assert.Equal(t, "error on update account", e.Error())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			s := tt.setup(ctrl)
			accID, limits := tt.args()
			tt.want(t, s.UpdateAccountLimits(context.Background(), accID, limits))
		})
	}
}

func TestService_UpdateAccountStatus(t *testing.T) {
	tests := []struct {
		name  string
		setup func(ctrl *gomock.Controller) *Service
		args  func() (int64, Status)
		want  func(t *testing.T, e error)
	}{
		{
			name: "when disable account with sucess",
			setup: func(ctrl *gomock.Controller) *Service {
				rep := NewMockRepository(ctrl)
				acc := buildAccount()
				rep.EXPECT().RetrieveAccountByID(gomock.Any(), int64(230513)).Return(acc, nil).Times(1)
				rep.EXPECT().UpdateExistingAccount(gomock.Any(), gomock.Any()).Return(nil).Times(1)
				return NewService(rep)
			},
			args: func() (int64, Status) {
				return int64(230513), Inative
			},
			want: func(t *testing.T, e error) {
				assert.Nil(t, e)
			},
		},
		{
			name: "when account not found",
			setup: func(ctrl *gomock.Controller) *Service {
				rep := NewMockRepository(ctrl)
				acc := Account{}
				rep.EXPECT().RetrieveAccountByID(gomock.Any(), int64(230513)).Return(acc, errors.New("account not found")).Times(1)
				return NewService(rep)
			},
			args: func() (int64, Status) {
				return int64(230513), Inative
			},
			want: func(t *testing.T, e error) {
				assert.NotNil(t, e)
				assert.Equal(t, "account not found", e.Error())
			},
		},
		{
			name: "when account update failed",
			setup: func(ctrl *gomock.Controller) *Service {
				rep := NewMockRepository(ctrl)
				acc := buildAccount()
				rep.EXPECT().RetrieveAccountByID(gomock.Any(), int64(230513)).Return(acc, nil).Times(1)
				rep.EXPECT().UpdateExistingAccount(gomock.Any(), gomock.Any()).Return(errors.New("account update failed")).Times(1)
				return NewService(rep)
			},
			args: func() (int64, Status) {
				return int64(230513), Inative
			},
			want: func(t *testing.T, e error) {
				assert.NotNil(t, e)
				assert.Equal(t, "account update failed", e.Error())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			s := tt.setup(ctrl)
			accID, status := tt.args()
			tt.want(t, s.UpdateAccountStatus(context.Background(), accID, status))
		})
	}
}
