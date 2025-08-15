package common

import (
	"sync"
)

type Aggregate interface {
	GetID() int64
	SetID(id string) *AggregateBase
	GetType() string
	SetType(typeName string)

	GetAggregateID() int64
	GetVersion() int64

	GetUncommittedEvents() []Event

	GetAppliedEvents() []Event
	SetAppliedEvents(events []Event) error

	When(event []Event) error
}

type AggregateBase struct {
	ID          string
	Type        string
	AggregateID int64
	Version     int64

	AppliedEvents     []Event //已经应用的事件列表
	UncommittedEvents []Event //尚未提交的事件列表

	When when

	Mu sync.RWMutex
}
type when func(event Event) error

func NewAggregateBase(when when) AggregateBase {
	return AggregateBase{
		AppliedEvents:     make([]Event, 0, 10),
		UncommittedEvents: make([]Event, 0, 10),
		Version:           0,
		When:              when,
	}
}

func (a *AggregateBase) GetID() string         { return a.ID }
func (a *AggregateBase) GetType() string       { return a.Type }
func (a *AggregateBase) GetAggregateID() int64 { return a.AggregateID }
func (a *AggregateBase) GetVersion() int64     { return a.Version }

func (a *AggregateBase) GetUncommittedEvents() []Event { return a.UncommittedEvents }
func (a *AggregateBase) ClearUncommittedEvents() {
	a.UncommittedEvents = make([]Event, 0, 10)
}

func (a *AggregateBase) GetAppliedEvents() []Event { return a.AppliedEvents }
func (a *AggregateBase) SetAppliedEvents(events []Event) {
	a.AppliedEvents = events
}
