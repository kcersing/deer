package aggregate

import "kcers-order/biz/infras/events"

type Aggregate interface {
	GetID() string
	GetType() string
	GetAggregateID() int64
	GetVersion() int64

	GetUncommittedEvents() []events.Event
	GetAppliedEvents() []events.Event
}

type AggregateBase struct {
	ID          string
	Type        string
	AggregateID int64
	Version     int64

	AppliedEvents     []events.Event //已经应用的事件列表
	UncommittedEvents []events.Event //尚未提交的事件列表

	when when
}
type when func(event events.Event) error

func NewAggregateBase(when when) AggregateBase {
	return AggregateBase{
		AppliedEvents:     make([]events.Event, 0, 10),
		UncommittedEvents: make([]events.Event, 0, 10),
		Version:           0,
		when:              when,
	}
}

func (a *AggregateBase) GetID() string         { return a.ID }
func (a *AggregateBase) GetType() string       { return a.Type }
func (a *AggregateBase) GetAggregateID() int64 { return a.AggregateID }
func (a *AggregateBase) GetVersion() int64     { return a.Version }

func (a *AggregateBase) GetUncommittedEvents() []events.Event { return a.UncommittedEvents }
func (a *AggregateBase) ClearUncommittedEvents() {
	a.UncommittedEvents = make([]events.Event, 0, 10)
}

func (a *AggregateBase) GetAppliedEvents() []events.Event { return a.AppliedEvents }
func (a *AggregateBase) SetAppliedEvents(events []events.Event) {
	a.AppliedEvents = events
}
