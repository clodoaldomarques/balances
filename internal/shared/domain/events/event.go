package events

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type EventType string

type Event struct {
	EventID   uuid.UUID `json:"event_id"`
	EventType EventType `json:"event_type`
	Data      any       `json:"data"`
	EventDate time.Time `json:"event_date"`
}

func (e Event) ToMessage() *string {
	evt, _ := json.Marshal(e)
	msg := string(evt)
	return &msg
}
