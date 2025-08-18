package common

import (
	"github.com/google/uuid"
	"time"
)

type Event interface {
	GetId() string
	GetType() string
	GetAggregateID() int64
	GetVersion() int64
	GetData() EventData

	SetType(eventType string)
	SetVersion(version int64)
	SetData(data EventData)
}
type EventBase struct {
	EventID       string
	EventType     string
	AggregateID   int64
	AggregateType string
	Version       int64
	Data          EventData
	Timestamp     time.Time
}

func NewBaseEvent(aggregate Aggregate, eventType string) EventBase {
	return EventBase{
		EventID:       uuid.New().String(),
		EventType:     eventType,
		AggregateID:   aggregate.GetID(),
		AggregateType: aggregate.GetType(),
		Version:       aggregate.GetVersion(),
		Timestamp:     time.Now(),
	}
}

func (e *EventBase) GetID() string            { return e.EventID }
func (e *EventBase) GetType() string          { return e.EventType }
func (e *EventBase) GetAggregateID() int64    { return e.AggregateID }
func (e *EventBase) GetAggregateType() string { return e.AggregateType }
func (e *EventBase) GetVersion() int64        { return e.Version }
func (e *EventBase) GetTimestamp() time.Time  { return e.Timestamp }
func (e *EventBase) GetData() EventData       { return e.Data }

func (e *EventBase) SetType(eventType string) {
	e.EventType = eventType
}
func (e *EventBase) SetVersion(version int64) {
	e.Version = version
}

func (e *EventBase) SetData(data EventData) {
	e.Data = data
}

type EventData struct {
	Type  string
	Event interface{}
}

func NewEventData(event Event) EventData {
	return EventData{
		Type:  event.GetType(),
		Event: event,
	}
}
