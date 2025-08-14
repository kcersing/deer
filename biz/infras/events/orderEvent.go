package events

import (
	"kcers-order/biz/infras/aggregate"
	"kcers-order/biz/infras/states"
)

// CreatedOrderEvent 创建订单事件
type CreatedOrderEvent struct {
	BaseEvent
	Sn          string
	TotalAmount float64
	Items       []aggregate.Item
	CreatedId   int64
}

func (e *CreatedOrderEvent) GetType() states.OrderStatus { return states.Created }

// PaidOrderEvent 支付订单事件
type PaidOrderEvent struct {
	BaseEvent
	PayedAmount float64
	PayMethod   string
	CreatedId   int64
}

func (e *PaidOrderEvent) GetType() states.OrderStatus { return states.Paid }

// CancelledOrderEvent 取消订单事件
type CancelledOrderEvent struct {
	BaseEvent
	Reason    string
	CreatedId int64
}

func (e *CancelledOrderEvent) GetType() states.OrderStatus { return states.Cancelled }

// ShippedOrderEvent 发货事件
type ShippedOrderEvent struct {
	BaseEvent
	CreatedId int64
}

func (e *ShippedOrderEvent) GetType() states.OrderStatus { return states.Shipped }

// CompletedOrderEvent 完成订单事件
type CompletedOrderEvent struct {
	BaseEvent
	CreatedId int64
}

func (e *CompletedOrderEvent) GetType() states.OrderStatus { return states.Completed }

// RefundedEvent 退款事件
type RefundedEvent struct {
	BaseEvent
	RefundedAmount float64
	CreatedId      int64
}

func (e *RefundedEvent) GetType() states.OrderStatus { return states.Refunded }
