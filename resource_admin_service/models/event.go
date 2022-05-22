package models

import(
	"encoding/json"
)

type Event struct {
	EventType string   `json:"event_type"`
	Resource  Resource `json:"resource"`
}

func NewEvent(eventType string, resource Resource) *Event {
	return &Event{
		EventType: eventType,
		Resource: resource,
	}
}

func (e *Event) Encode() ([]byte, error) {
	return json.Marshal(e)
}
