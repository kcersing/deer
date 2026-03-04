package aggregate

import (
	"gen/kitex_gen/base"
	"order/biz/infras/common"
	"order/biz/infras/events"
	"time"
)

const (
	OrderAggregateType = "order"
)

// Order 聚合根
type Order struct {
	base.Order
	stateMachine *StateMachine
	common.AggregateBase
}

func NewOrder() (order *Order) {
	order = &Order{}
	// 优化：直接让 Order 实现 Whener 接口，而不是传递函数
	aggregateBase := common.NewAggregateBase(order)
	aggregateBase.SetAggregateType(OrderAggregateType)
	order.AggregateBase = *aggregateBase
	order.stateMachine = NewStateMachine(order)

	return order
}

// GetStateMachine 获取状态机
func (o *Order) GetStateMachine() *StateMachine {
	return o.stateMachine
}

func (o *Order) When(evt common.Event) error {
	switch e := evt.(type) {
	case *events.CreatedOrderEvent:
		return o.onCreated(e)
	case *events.CancelledOrderEvent:
		return o.onCancelled(e)
	case *events.CompletedOrderEvent:
		return o.onCompleted(e)
	case *events.PaidOrderEvent:
		return o.onPaid(e)
	case *events.RefundedOrderEvent:
		return o.onRefunded(e)
	case *events.ShippedOrderEvent:
		return o.onShipped(e)
	default:
		return common.ErrInvalidType
	}
}

func (o *Order) onCreated(eventData *events.CreatedOrderEvent) (err error) {
	o.Sn = eventData.Sn
	o.MemberId = eventData.MemberId
	o.CreatedId = eventData.CreatedId
	o.Items = eventData.Items
	o.TotalAmount = eventData.TotalAmount
	o.Status = eventData.GetType()
	return nil
}
func (o *Order) onCancelled(eventData *events.CancelledOrderEvent) (err error) {
	o.CancelledReason = eventData.Reason
	o.CreatedId = eventData.CreatedId
	o.CloseAt = eventData.Timestamp.Format(time.DateTime)
	o.Status = eventData.GetType()
	return nil
}
func (o *Order) onCompleted(eventData *events.CompletedOrderEvent) (err error) {
	o.CreatedId = eventData.CreatedId
	o.CompletionAt = eventData.Timestamp.Format(time.DateTime)
	o.Status = eventData.GetType()
	return nil
}
func (o *Order) onPaid(eventData *events.PaidOrderEvent) (err error) {
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
	o.Status = eventData.GetType()
	return nil
}
func (o *Order) onRefunded(eventData *events.RefundedOrderEvent) (err error) {
	o.OrderRefund.CreatedId = eventData.CreatedId
	o.OrderRefund.RefundAmount = eventData.RefundedAmount
	o.OrderRefund.RefundReason = eventData.Reason
	o.OrderRefund.RefundAt = eventData.Timestamp.Format(time.DateTime)
	o.Status = eventData.GetType()
	return nil
}
func (o *Order) onShipped(eventData *events.ShippedOrderEvent) (err error) {
	o.CreatedId = eventData.CreatedId
	o.Status = eventData.GetType()
	return nil
}
