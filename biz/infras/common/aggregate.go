package common

type Aggregate interface {
	GetAggregateType() string
	GetAggregateID() int64
	GetVersion() int64

	GetUncommittedEvents() []Event
	ClearUncommittedEvents()
	GetAppliedEvents() []Event
	SetAppliedEvents(events []Event)
	SetAggregateType(aggregateType string)
	When
	Apply
	Load
}
type When interface {
	When(event Event) error
}
type Apply interface {
	Apply(event Event) error
}
type Load interface {
	Load(events []Event) error
}

var _ Aggregate = &AggregateBase{}

type AggregateBase struct {
	AggregateID   int64
	AggregateType string

	Version int64

	appliedEvents     []Event //已经应用的事件列表
	uncommittedEvents []Event //尚未提交的事件列表

	when when
}

func (a *AggregateBase) When(event Event) error {
	return a.when(event)
}

func NewAggregateBase(when when) *AggregateBase {
	if when == nil {
		return nil
	}

	return &AggregateBase{
		Version:           0,
		appliedEvents:     []Event{},
		uncommittedEvents: []Event{},
		when:              when,
	}
}

type when func(event Event) error

func (a *AggregateBase) GetAggregateType() string { return a.AggregateType }
func (a *AggregateBase) GetAggregateID() int64    { return a.AggregateID }
func (a *AggregateBase) GetVersion() int64        { return a.Version }

func (a *AggregateBase) GetUncommittedEvents() []Event {
	es := make([]Event, len(a.uncommittedEvents))
	copy(es, a.uncommittedEvents)
	return es
}
func (a *AggregateBase) ClearUncommittedEvents() {
	a.uncommittedEvents = []Event{}
}

func (a *AggregateBase) GetAppliedEvents() []Event { return a.appliedEvents }
func (a *AggregateBase) SetAppliedEvents(events []Event) {
	a.appliedEvents = events
}
func (a *AggregateBase) SetAggregateType(aggregateType string) {
	a.AggregateType = aggregateType
}
func (a *AggregateBase) Load(events []Event) error {
	for _, evt := range events {
		if evt.GetAggregateID() != a.AggregateID {
			return ErrInvalidAggregate
		}
		if err := a.when(evt); err != nil {
			return err
		}
		a.Version++
	}
	return nil
}

func (a *AggregateBase) Apply(event Event) error {

	if event.GetAggregateID() != a.GetAggregateID() {
		return ErrInvalidAggregateID
	}

	event.SetType(a.GetAggregateType())
	if err := a.when(event); err != nil {
		return err
	}
	a.Version++
	event.SetVersion(a.GetVersion())
	a.uncommittedEvents = append(a.uncommittedEvents, event)
	return nil
}

type Item struct {
	ProductId int64
	Quantity  int64
	Price     float64
	Name      string
}
