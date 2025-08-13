package events

import "time"

type Event interface {
}
type BaseEvent struct {
	EventID     string
	AggregateID int64
	Version     int64
	Type        string
	Timestamp   time.Time
}

type EventData struct {
	Type  string
	Event interface{}
}
