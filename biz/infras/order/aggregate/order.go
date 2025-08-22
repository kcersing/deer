package aggregate

import (
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/order/events"
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

	Nature string
	mu     sync.RWMutex
	common.AggregateBase
}

func NewOrder() (order *Order) {
	order = &Order{}
	base := common.NewAggregateBase(order.When)
	base.SetAggregateType(OrderAggregateType)
	order.AggregateBase = *base
	order.stateMachine = NewStateMachine(order)

	return order
}

func (o *Order) When(evt common.Event) error {

	switch evt.GetType() {
	case string(status.Created):
		return o.onCreated(evt)
	case string(status.Cancelled):
		return o.onCancelled(evt)
	case string(status.Completed):
		return o.onCompleted(evt)
	case string(status.Paid):
		return o.onPaid(evt)
	case string(status.Refunded):
		return o.onRefunded(evt)
	case string(status.Shipped):
		return o.onShipped(evt)
	default:
		return common.ErrInvalidType
	}
	//}
}

func (o *Order) onCreated(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if eventData, ok := evt.(*events.CreatedOrderEvent); ok {
		o.Sn = eventData.Sn
		o.MemberId = eventData.MemberId
		o.CreatedId = eventData.CreatedId
		o.Items = eventData.Items
		o.TotalAmount = eventData.TotalAmount
		o.Status = status.OrderStatus(eventData.GetType())
	}

	return nil
}
func (o *Order) onCancelled(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if eventData, ok := evt.(*events.CancelledOrderEvent); ok {

		//o.Reason   = eventData.Reason
		o.CreatedId = eventData.CreatedId

	}
	return nil
}
func (o *Order) onCompleted(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if eventData, ok := evt.(*events.CompletedOrderEvent); ok {

		o.CreatedId = eventData.CreatedId

	}
	return nil
}
func (o *Order) onPaid(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if eventData, ok := evt.(*events.PaidOrderEvent); ok {

		o.CreatedId = eventData.CreatedId

	}
	return nil
}
func (o *Order) onRefunded(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if eventData, ok := evt.(*events.RefundedOrderEvent); ok {

		o.CreatedId = eventData.CreatedId
	}
	return nil
}
func (o *Order) onShipped(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if eventData, ok := evt.(*events.ShippedOrderEvent); ok {

		o.CreatedId = eventData.CreatedId
	}
	return nil
}
