package common

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
}

type AggregateBase struct {
	AggregateID   int64
	AggregateType string
	Version       int64

	AppliedEvents     []Event //已经应用的事件列表
	UncommittedEvents []Event //尚未提交的事件列表

	when when
}

type when func(event Event) error

func (a *AggregateBase) GetAggregateType() string { return a.AggregateType }
func (a *AggregateBase) GetAggregateID() int64    { return a.AggregateID }
func (a *AggregateBase) GetVersion() int64        { return a.Version }

func (a *AggregateBase) GetUncommittedEvents() []Event { return a.UncommittedEvents }
func (a *AggregateBase) ClearUncommittedEvents() {
	a.UncommittedEvents = make([]Event, 0, 10)
}

func (a *AggregateBase) GetAppliedEvents() []Event { return a.AppliedEvents }
func (a *AggregateBase) SetAppliedEvents(events []Event) {
	a.AppliedEvents = events
}

//func (a *AggregateBase) When(events []Event) {
//	for _, event := range events {
//		err := a.when(event)
//		if err != nil {
//			return
//		}
//	}
//
//	a.AppliedEvents = events
//	a.Version++
//}

type Item struct {
	ProductId int64
	Quantity  int64
	Price     float64
	Name      string
}
