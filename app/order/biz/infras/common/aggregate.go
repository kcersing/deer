package common

// Whener 接口定义了应用事件的核心逻辑
type Whener interface {
	When(event Event) error
}

type Aggregate interface {
	GetAggregateType() string
	GetAggregateID() int64
	GetVersion() int64

	GetUncommittedEvents() []Event
	ClearUncommittedEvents()
	GetAppliedEvents() []Event
	SetAppliedEvents(events []Event)
	SetAggregateID(aggregateID int64)
	SetAggregateType(aggregateType string)

	Apply(event Event) error
	Load(events []Event) error
	Whener
}

var _ Aggregate = &AggregateBase{}

type AggregateBase struct {
	AggregateID   int64
	AggregateType string
	Version       int64

	appliedEvents     []Event // 已经应用的事件列表
	uncommittedEvents []Event // 尚未提交的事件列表

	whener Whener // 聚合根自身，实现了 When 方法
}

func NewAggregateBase(whener Whener) *AggregateBase {
	if whener == nil {
		// 当聚合根未实现 Whener 接口时，这是个编程错误
		panic("whener cannot be nil")
	}

	return &AggregateBase{
		Version:           0,
		appliedEvents:     []Event{},
		uncommittedEvents: []Event{},
		whener:            whener,
	}
}

func (a *AggregateBase) GetAggregateType() string { return a.AggregateType }
func (a *AggregateBase) GetAggregateID() int64    { return a.AggregateID }
func (a *AggregateBase) GetVersion() int64        { return a.Version }

func (a *AggregateBase) GetUncommittedEvents() []Event {
	// 返回一个副本以防止外部修改
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
func (a *AggregateBase) SetAggregateID(aggregateID int64) {
	a.AggregateID = aggregateID
}
func (a *AggregateBase) SetAggregateType(aggregateType string) {
	a.AggregateType = aggregateType
}

// Load 从历史事件中重建聚合状态
func (a *AggregateBase) Load(events []Event) error {
	for _, evt := range events {
		if a.AggregateID != 0 && evt.GetAggregateID() != a.AggregateID {
			return ErrInvalidAggregate
		}
		if err := a.whener.When(evt); err != nil {
			return err
		}
		a.Version = evt.GetVersion() // 从事件中恢复版本号
	}
	a.SetAppliedEvents(events)
	return nil
}

// Apply 应用一个新的事件
func (a *AggregateBase) Apply(event Event) error {
	// 在应用第一个事件（通常是创建事件）之前，聚合ID可能为0
	if a.AggregateID != 0 && event.GetAggregateID() != a.GetAggregateID() {
		return ErrInvalidAggregateID
	}

	// 调用聚合根自己的 When 方法来更新状态
	if err := a.whener.When(event); err != nil {
		return err
	}

	// 递增版本号并添加到未提交事件列表
	a.Version++
	event.SetVersion(a.GetVersion())
	a.uncommittedEvents = append(a.uncommittedEvents, event)
	return nil
}

// When 是 AggregateBase 对 Whener 接口的默认实现，
// 但它实际上会委托给聚合根（如 Order）的具体实现。
// 这使得 AggregateBase 自身也满足了 Aggregate 接口。
func (a *AggregateBase) When(event Event) error {
	return a.whener.When(event)
}
