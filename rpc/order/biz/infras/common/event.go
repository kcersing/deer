package common

import (
	"time"
)

type Event interface {
	GetId() string
	GetType() string
	GetAggregateID() int64
	GetVersion() int64
	GetAggregateType() string
	GetTimestamp() time.Time

	SetAggregateID(aggregateID int64)
	SetAggregateType(aggregateType string)
	SetType(eventType string)
	SetVersion(version int64)
}
type EventBase struct {
	EventID       string
	EventType     string
	AggregateID   int64
	AggregateType string
	Version       int64
	Timestamp     time.Time
}

func (e *EventBase) GetId() string {
	return e.EventID
}
func (e *EventBase) GetType() string {
	return e.EventType
}
func (e *EventBase) GetAggregateID() int64 {
	return e.AggregateID
}
func (e *EventBase) GetAggregateType() string {
	return e.AggregateType
}
func (e *EventBase) GetVersion() int64 {
	return e.Version
}
func (e *EventBase) GetTimestamp() time.Time {
	return e.Timestamp
}

func (e *EventBase) SetAggregateID(aggregateID int64) {
	e.AggregateID = aggregateID
}
func (e *EventBase) SetAggregateType(aggregateType string) {
	e.AggregateType = aggregateType
}
func (e *EventBase) SetType(eventType string) {
	e.EventType = eventType
}
func (e *EventBase) SetVersion(version int64) {
	e.Version = version
}

type EventData struct {
	Event Event
}
