package aggregate

import (
	"gen/kitex_gen/base"

	"order/biz/infras/common"
	"order/biz/infras/events"
	"time"

	"sync"
)

const (
	OrderAggregateType = "order"
)

type Order struct {
	base.Order

	stateMachine *StateMachine
	common.AggregateBase
	mu sync.RWMutex
}

func NewOrder() (order *Order) {
	order = &Order{}
	aggregateBase := common.NewAggregateBase(order.When)
	aggregateBase.SetAggregateType(OrderAggregateType)
	order.AggregateBase = *aggregateBase
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
		o.CancelledReason = eventData.Reason
		o.CreatedId = eventData.CreatedId
		o.CloseAt = eventData.Timestamp.Format(time.DateTime)
	}
	return nil
}
func (o *Order) onCompleted(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if eventData, ok := evt.(*events.CompletedOrderEvent); ok {

		o.CreatedId = eventData.CreatedId
		o.CompletionAt = eventData.Timestamp.Format(time.DateTime)

	}
	return nil
}
func (o *Order) onPaid(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if eventData, ok := evt.(*events.PaidOrderEvent); ok {
		var orderPay base.OrderPay
		orderPay.CreatedId = eventData.CreatedId
		orderPay.Pay = eventData.PayedAmount
		orderPay.PayWay = eventData.PayMethod
		orderPay.PayAt = eventData.Timestamp.Format(time.DateTime)
		orderPay.Remission = eventData.Remission
		orderPay.Reason = eventData.Reason
		orderPay.PaySn = eventData.PaySn
		orderPay.PrepayId = eventData.PrepayId
		orderPay.PayExtra = eventData.PayExtra
		o.OrderPays = append(o.OrderPays, &orderPay)
	}
	return nil
}
func (o *Order) onRefunded(evt common.Event) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if eventData, ok := evt.(*events.RefundedOrderEvent); ok {
		o.OrderRefund.CreatedId = eventData.CreatedId
		o.OrderRefund.RefundAmount = eventData.RefundedAmount
		o.OrderRefund.RefundReason = eventData.Reason
		o.OrderRefund.RefundAt = eventData.Timestamp.Format(time.DateTime)
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
