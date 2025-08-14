package events

import (
	"github.com/google/uuid"
	"kcers-order/biz/infras/aggregate"
	"time"
)

type Event interface {
	GetId() string
	GetType() string
	GetAggregateID() int64
	GetVersion() int64
}
type EventBase struct {
	EventID       string
	EventType     string
	AggregateID   int64
	AggregateType string
	Version       int64

	Timestamp time.Time
}

func NewBaseEvent(aggregate aggregate.Aggregate, eventType string) EventBase {
	return EventBase{
		EventID:       uuid.New().String(),
		EventType:     eventType,
		AggregateID:   aggregate.GetID(),
		AggregateType: aggregate.GetType(),
		Version:       aggregate.GetVersion(),
		Timestamp:     time.Now(),
	}
}

func (e *EventBase) GetID() string           { return e.EventID }
func (e *EventBase) GetType() string         { return e.EventType }
func (e *EventBase) GetAggregateID() int64   { return e.AggregateID }
func (e *EventBase) GetVersion() int64       { return e.Version }
func (e *EventBase) GetTimestamp() time.Time { return e.Timestamp }

type EventData struct {
	Type  string
	Event interface{}
}
