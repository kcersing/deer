package aggregate

import (
	"errors"
	"gen/kitex_gen/base"
	"gen/kitex_gen/order"

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
	case *events.PayingOrderEvent:
		return o.onPaying(e)
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
func (o *Order) onPaying(e *events.PayingOrderEvent) (err error) {
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
	o.Actual += e.Amount
	o.Remission += e.Remission
	return nil
}
func (o *Order) onPaid(e *events.PaidOrderEvent) (err error) {
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
	o.Status = e.GetType()
	return nil
}

func (o *Order) Create(req *order.CreateOrderReq) error {
	event := events.NewCreatedOrderEvent(req)
	err := o.Apply(event)
	if err != nil {
		return err
	}
	return err
}

func (o *Order) Paying(req *order.PaymentReq) error {
	if err := o.stateMachine.ValidateTransition(common.Paying); err != nil {
		return err
	}
	if req.GetAmount() <= 0 {
		return errors.New("支付金额必须为正数")
	}
	if req.GetAmount() > o.TotalAmount-o.Actual-o.Remission {
		return errors.New("支付金额超出订单待付金额")
	}
	event := events.NewPayingOrderEvent(req)
	err := o.Apply(event)
	if err != nil {
		return err
	}
	if o.TotalAmount-o.Actual-o.Remission == 0 {
		err = o.Paid()
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *Order) Paid() error {
	if err := o.stateMachine.ValidateTransition(common.Paid); err != nil {
		return err
	}
	event := events.NewPaidOrderEvent(o.AggregateID)
	err := o.Apply(event)
	if err != nil {
		return err
	}
	return err
}
func (o *Order) Shipped() error {
	if err := o.stateMachine.ValidateTransition(common.Shipped); err != nil {
		return err
	}
	event := events.NewShippedOrderEvent(o.AggregateID)
	event.MemberId = o.MemberId
	event.OrderId = o.Id
	event.UserId = o.CreatedId
	event.Items = o.Items
	err := o.Apply(event)
	if err != nil {
		return err
	}
	return err
}
func (o *Order) Cancel(req *order.CancelledOrderReq) error {
	// 1. 业务规则封装在内部
	if err := o.stateMachine.ValidateTransition(common.Cancelled); err != nil {

		return err
	}

	// if o.IsSpecialProduct() { return errors.New("特殊商品不可取消") }

	event := events.NewCancelledOrderEvent(req)

	err := o.Apply(event)
	if err != nil {
		return err
	}
	return nil
}
func (o *Order) Refund(req *order.RefundOrderReq) error {
	if err := o.stateMachine.ValidateTransition(common.Refunded); err != nil {
		return err
	}
	event := events.NewRefundedOrderEvent(req)
	err := o.Apply(event)
	if err != nil {
		return err
	}

	return nil
}
