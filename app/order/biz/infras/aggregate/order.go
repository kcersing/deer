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

func (o *Order) onCreated(e *events.CreatedOrderEvent) (err error) {
	o.Sn = e.Sn
	o.MemberId = e.MemberId
	o.CreatedId = e.CreatedId
	o.Items = e.Items
	o.TotalAmount = e.TotalAmount
	o.Status = e.GetType()
	return nil
}
func (o *Order) onCancelled(e *events.CancelledOrderEvent) (err error) {

	o.CancelledReason = e.Reason
	o.CreatedId = e.CreatedId
	o.CloseAt = e.Timestamp.Format(time.DateTime)
	o.Status = e.GetType()
	return nil
}
func (o *Order) onCompleted(e *events.CompletedOrderEvent) (err error) {
	o.CreatedId = e.CreatedId
	o.CompletionAt = e.Timestamp.Format(time.DateTime)
	o.Status = e.GetType()
	return nil
}
func (o *Order) onPaid(e *events.PaidOrderEvent) (err error) {

	// 可以在这里添加其他业务规则检查
	// 例如：检查支付金额是否正确等
	//if payedAmount <= 0 {
	//	return errors.New("支付金额必须为正数")
	//}
	var orderPay base.OrderPay
	orderPay.CreatedId = e.CreatedId
	orderPay.Pay = e.Amount
	orderPay.PayWay = e.Method
	orderPay.PayAt = e.Timestamp.Format(time.DateTime)
	orderPay.Remission = e.Remission
	orderPay.Reason = e.Reason
	orderPay.PaySn = e.PaySn
	orderPay.PrepayId = e.PrepayId
	orderPay.PayExtra = e.PayExtra
	o.OrderPays = append(o.OrderPays, &orderPay)
	o.Status = e.GetType()
	return nil
}
func (o *Order) onRefunded(e *events.RefundedOrderEvent) (err error) {
	o.OrderRefund.CreatedId = e.CreatedId
	o.OrderRefund.Amount = e.Amount
	o.OrderRefund.Reason = e.Reason
	o.OrderRefund.RefundAt = e.Timestamp.Format(time.DateTime)
	o.Status = e.GetType()
	return nil
}
func (o *Order) onShipped(e *events.ShippedOrderEvent) (err error) {
	o.CreatedId = e.CreatedId
	o.Status = e.GetType()
	return nil
}

func (o *Order) Cancel(userID int64, reason string) error {
	// 1. 业务规则封装在内部
	if err := o.stateMachine.ValidateTransition(common.Cancelled); err != nil {
		return err
	}

	// if o.IsSpecialProduct() { return errors.New("特殊商品不可取消") }

	cancelEvent := events.NewCancelledOrderEvent(o.GetAggregateID(), reason, userID)

	return o.Apply(cancelEvent)
}
