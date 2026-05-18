package accounts

import (
	"time"

	"github.com/clodoaldomarques/balances/internal/shared/commons"
	"github.com/clodoaldomarques/balances/internal/shared/domain/events"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

var (
	CREATE_ACCOUNT events.Type = "create_account"
	UPDATE_ACCOUNT events.Type = "update_account"
	PROCESS_ENTRY  events.Type = "process_entry"
)

type CreateAccountEvent struct {
	AccountID int64              `json:"account_id"`
	OrgID     string             `json:"org_id"`
	Limits    commons.DecimalMap `json:"limits"`
	Balances  commons.DecimalMap `json:"balances"`
	CreatedAt time.Time          `json:"created_at"`
	Status    string             `json:"status"`
	Version   int64              `json:"version"`
}

func buildCreateAccountEvent(a Account) events.Event {
	evt := CreateAccountEvent{
		AccountID: a.AccountID,
		OrgID:     a.OrgID,
		Limits:    a.Limits,
		Balances:  a.Balances,
		CreatedAt: a.CreatedAt,
		Status:    string(a.Status),
		Version:   a.Version,
	}

	return events.Event{
		EventID:   uuid.New(),
		EventType: CREATE_ACCOUNT,
		Data:      evt,
		EventDate: time.Now(),
	}
}

type UpdateAccountEvent struct {
	AccountID int64              `json:"account_id"`
	OrgID     string             `json:"org_id"`
	Limits    commons.DecimalMap `json:"limits"`
	Balances  commons.DecimalMap `json:"balances"`
	Status    string             `json:"status"`
	UpdatedAt time.Time          `json:"updated_at"`
	Version   int64              `json:"version"`
}

func buildUpdateAccountEvent(a Account) events.Event {
	evt := UpdateAccountEvent{
		AccountID: a.AccountID,
		OrgID:     a.OrgID,
		Limits:    a.Limits,
		Balances:  a.Balances,
		UpdatedAt: a.UpdatedAt,
		Status:    string(a.Status),
		Version:   a.Version,
	}

	return events.Event{
		EventID:   uuid.New(),
		EventType: UPDATE_ACCOUNT,
		Data:      evt,
		EventDate: time.Now(),
	}
}

type ProcessEntryEvent struct {
	AccountID  int64              `json:"account_id"`
	OrgID      string             `json:"org_id"`
	TrackingID string             `json:"tracking_id"`
	Impacts    []ImpactEvent      `json:"impacts"`
	Limits     commons.DecimalMap `json:"limits"`
	Balances   commons.DecimalMap `json:"balances"`
	Version    int64              `json:"version"`
	CreatedAt  time.Time          `json:"created_at"`
}

func buildProcessEntryEvent(a Account, e Entry) events.Event {
	evt := ProcessEntryEvent{
		AccountID:  a.AccountID,
		OrgID:      a.OrgID,
		TrackingID: e.TrackingID,
		Impacts:    buildImpactEvents(e.Impacts),
		Limits:     a.Limits,
		Balances:   a.Balances,
		Version:    a.Version,
		CreatedAt:  e.CreatedAt,
	}

	return events.Event{
		EventID:   uuid.New(),
		EventType: PROCESS_ENTRY,
		Data:      evt,
		EventDate: time.Now(),
	}
}

type ImpactEvent struct {
	Balance   string          `json:"balance"`
	Operation string          `json:"operation"`
	Amount    decimal.Decimal `json:"amount"`
	Rules     []string        `json:"rules,omitempty"`
}

func buildImpactEvents(impacts []Impact) []ImpactEvent {
	evts := make([]ImpactEvent, 0, len(impacts))
	for _, i := range impacts {
		new := ImpactEvent{
			Balance:   i.Balance,
			Operation: i.Operation,
			Amount:    i.Amount,
			Rules:     i.Rules,
		}
		evts = append(evts, new)
	}
	return evts
}
