package aggregate

import (
	"deer/kitex_gen/deer/order"
	"deer/rpc/order/biz/infras/common"
	"deer/rpc/order/biz/infras/events"

	"sync"
)

const (
	OrderAggregateType = "order"
)

type Order struct {
	order.Order
	stateMachine *StateMachine
	mu           sync.RWMutex
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
	case string(common.Created):
		return o.onCreated(evt)
	case string(common.Cancelled):
		return o.onCancelled(evt)
	case string(common.Completed):
		return o.onCompleted(evt)
	case string(common.Paid):
		return o.onPaid(evt)
	case string(common.Refunded):
		return o.onRefunded(evt)
	case string(common.Shipped):
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
		o.Status = eventData.GetType()
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
