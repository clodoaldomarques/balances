package events

import (
	"time"

	"github.com/shopspring/decimal"
)

type Impact struct {
	Balance       string          `json:"balance"`
	Amount        decimal.Decimal `json:"amount"`
	OperationType string          `json:"operation_type"`
}

type CreateAccountEvent struct {
	TrackingID string                     `json:"tracking_id"`
	EventType  string                     `json:"event_type"`
	AccountID  int64                      `json:"account_id"`
	TenantID   string                     `json:"tenant_id"`
	Limits     map[string]decimal.Decimal `json:"limits"`
	Balances   map[string]decimal.Decimal `json:"balances"`
	Status     string                     `json:"status,omitempty"`
	EventDate  time.Time                  `json:"event_date"`
	Version    int64                      `json:"version"`
}

func (c CreateAccountEvent) GetTrackingID() string {
	return c.TrackingID
}
func (c CreateAccountEvent) GetEventType() string {
	return c.EventType
}
func (c CreateAccountEvent) GetAccountID() int64 {
	return c.AccountID
}
func (c CreateAccountEvent) GetTenantID() string {
	return c.TenantID
}
func (c CreateAccountEvent) GetEventDate() time.Time {
	return c.EventDate
}
func (c CreateAccountEvent) GetVersion() int64 {
	return c.Version
}
func (c CreateAccountEvent) GetEventData() string {
	return ""
}

type ChangeBalanceEvent struct {
	TrackingID string                     `json:"tracking_id"`
	EventType  string                     `json:"event_type"`
	AccountID  int64                      `json:"account_id"`
	TenantID   string                     `json:"tenant_id"`
	Impacts    []Impact                   `json:"impacts"`
	Balances   map[string]decimal.Decimal `json:"balances"`
	EventDate  time.Time                  `json:"event_date"`
	Version    int64                      `json:"version"`
}

func (c ChangeBalanceEvent) GetTrackingID() string {
	return c.TrackingID
}
func (c ChangeBalanceEvent) GetEventType() string {
	return c.EventType
}
func (c ChangeBalanceEvent) GetAccountID() int64 {
	return c.AccountID
}
func (c ChangeBalanceEvent) GetTenantID() string {
	return c.TenantID
}
func (c ChangeBalanceEvent) GetEventDate() time.Time {
	return c.EventDate
}
func (c ChangeBalanceEvent) GetVersion() int64 {
	return c.Version
}
func (c ChangeBalanceEvent) GetEventData() string {
	return c.EventDate.String()
}

type ChangeLimitsEvent struct {
	TrackingID string                     `json:"tracking_id"`
	EventType  string                     `json:"event_type"`
	AccountID  int64                      `json:"account_id"`
	TenantID   string                     `json:"tenant_id"`
	Limits     map[string]decimal.Decimal `json:"limits"`
	EventDate  time.Time                  `json:"event_date"`
	Version    int64                      `json:"version"`
}

func (c ChangeLimitsEvent) GetTrackingID() string {

}
func (c ChangeLimitsEvent) GetEventType() string {

}
func (c ChangeLimitsEvent) GetAccountID() int64 {

}
func (c ChangeLimitsEvent) GetTenantID() string {

}
func (c ChangeLimitsEvent) GetEventDate() time.Time {

}
func (c ChangeLimitsEvent) GetVersion() int64 {

}
func (c ChangeLimitsEvent) GetEventData() string {

}

type ChangeStatusEvent struct {
	TrackingID string    `json:"tracking_id"`
	EventType  string    `json:"event_type"`
	AccountID  int64     `json:"account_id"`
	TenantID   string    `json:"tenant_id"`
	Status     string    `json:"status"`
	EventDate  time.Time `json:"event_date"`
	Version    int64     `json:"version"`
}

func (c ChangeStatusEvent) GetTrackingID() string {

}
func (c ChangeStatusEvent) GetEventType() string {

}
func (c ChangeStatusEvent) GetAccountID() int64 {

}
func (c ChangeStatusEvent) GetTenantID() string {

}
func (c ChangeStatusEvent) GetEventDate() time.Time {

}
func (c ChangeStatusEvent) GetVersion() int64 {

}
func (c ChangeStatusEvent) GetEventData() string {

}
