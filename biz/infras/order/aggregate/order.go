package aggregate

import (
	"fmt"
	"github.com/pkg/errors"
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/status"
	"sync"
)

const (
	OrderAggregateType = "order"
)

type Order struct {
	Sn string

	MemberId  int64
	CreatedId int64

	Items       []common.Item
	TotalAmount float64

	Status       status.OrderStatus
	stateMachine *StateMachine

	mu sync.RWMutex

	common.AggregateBase
}

func NewOrder(sn string, items []common.Item, amount float64, memberId int64, userId int64) (order *Order) {
	order = &Order{
		Sn:          sn,
		MemberId:    memberId,
		CreatedId:   userId,
		Items:       items,
		TotalAmount: amount,
		Status:      status.Created,
	}
	order.AggregateType = OrderAggregateType
	order.stateMachine = NewStateMachine(order)

	return order
}

//func (o *Order) when(evt common.Event) error {
//
//	switch evt.GetType() {
//	case string(status.Created):
//		return o.onCreated(evt)
//	case string(status.Cancelled):
//		return o.onCancelled(evt)
//	case string(status.Completed):
//		return o.onCompleted(evt)
//	case string(status.Paid):
//		return o.onPaid(evt)
//	case string(status.Refunded):
//		return o.onRefunded(evt)
//	case string(status.Shipped):
//		return o.onShipped(evt)
//	default:
//		return common.ErrInvalidType
//	}
//}

func (o *Order) AddUncommittedEvent(event common.Event) error {
	o.mu.Lock()
	defer o.mu.Unlock()
	if event.GetAggregateID() != o.AggregateID {
		return errors.New("aggregate id not match")
	}

	//if err := o.When(event); err != nil {
	//	return err
	//}
	o.Version++
	//event.Version = o.GetVersion()

	o.UncommittedEvents = append(o.UncommittedEvents, event)
	return nil
}
func (o *Order) GetUncommittedEvents() []common.Event {
	o.mu.RLock()
	defer o.mu.RUnlock()
	es := make([]common.Event, len(o.UncommittedEvents))
	copy(es, o.UncommittedEvents)
	return es
}
func (o *Order) ClearUncommittedEvents() {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.UncommittedEvents = []common.Event{}
}

// StateMachine 订单状态机
type StateMachine struct {
	order *Order
}

func NewStateMachine(order *Order) *StateMachine {
	return &StateMachine{order: order}
}

// 定义状态转换规则
var transitions = map[status.OrderStatus][]status.OrderStatus{
	status.Created: {status.Paid, status.Cancelled},
	status.Paid:    {status.Shipped, status.Refunded, status.Cancelled},
	status.Shipped: {status.Completed, status.Refunded},
}

// Transition 执行状态转换
func (m *StateMachine) Transition(target status.OrderStatus, event common.Event) error {
	m.order.mu.Lock()
	current := m.order.Status
	m.order.mu.Unlock()

	if current == target {
		return nil // 状态未变更
	}

	// 检查是否允许转换
	allowed := false
	for _, status := range transitions[current] {
		if status == target {
			allowed = true
			break
		}
	}
	if !allowed {
		return fmt.Errorf("状态转换不允许: %s -> %s", current, target)
	}

	// 执行转换
	m.order.mu.Lock()
	m.order.Status = target
	m.order.Version++
	m.order.UncommittedEvents = append(m.order.UncommittedEvents, event)
	m.order.mu.Unlock()

	return nil
}

func (o *Order) onCreated(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	//if eventData, ok := evt.GetData().Event.(*events.CreatedOrderEvent); ok {
	//	o.Id = evt.GetAggregateID()
	//	o.Items = eventData.Items
	//	o.TotalAmount = eventData.TotalAmount
	//	o.MemberId = eventData.MemberId
	//	o.CreatedId = eventData.CreatedId
	//}

	return nil
}
func (o *Order) onCancelled(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	//if eventData, ok := evt.GetData().Event.(*events.CancelledOrderEvent); ok {
	//
	//	klog.Infof("onCancelled: %v", eventData)
	//
	//}
	return nil
}
func (o *Order) onCompleted(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	//if eventData, ok := evt.GetData().Event.(*events.CompletedOrderEvent); ok {
	//
	//	klog.Infof("onCompleted: %v", eventData)
	//
	//}
	return nil
}
func (o *Order) onPaid(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	//if eventData, ok := evt.GetData().Event.(*events.PaidOrderEvent); ok {
	//
	//	klog.Infof("onPaid: %v", eventData)
	//
	//}
	return nil
}
func (o *Order) onRefunded(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	//if eventData, ok := evt.GetData().Event.(*events.RefundedOrderEvent); ok {
	//
	//	klog.Infof("onRefunded: %v", eventData)
	//
	//}
	return nil
}
func (o *Order) onShipped(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	//if eventData, ok := evt.GetData().Event.(*events.ShippedOrderEvent); ok {
	//
	//	klog.Infof("onShipped: %v", eventData)
	//
	//}
	return nil
}
