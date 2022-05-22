package models

import (
	"encoding/json"
)

const (
	CreateEvent = "CREATE"
	UpdateEvent = "UPDATE"
	DeleteEvent = "DELETE"
)

type Event struct {
	EventType string   `json:"event_type"`
	Resource  Resource `json:"resource"`
}

func DecodeEvent(eventjson []byte) (Event, error) {
	event := Event{}
	err := json.Unmarshal(eventjson, &event)
	if err != nil {
		return event, err
	}
	return event, nil
}
