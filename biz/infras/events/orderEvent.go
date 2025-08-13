package events

// CreatedOrderEvent 创建订单事件
type CreatedOrderEvent struct {
	BaseEvent
	Sn        string
	CreatedId int64
}

// PaidOrderEvent 支付订单事件
type PaidOrderEvent struct {
	BaseEvent
	PayedAmount float64
	PayMethod   string
}

// CancelledOrderEvent 取消订单事件
type CancelledOrderEvent struct {
	BaseEvent
	Reason string
}

// ShippedOrderEvent 发货事件
type ShippedOrderEvent struct {
	BaseEvent
}

// CompletedOrderEvent 完成订单事件
type CompletedOrderEvent struct {
	BaseEvent
}

// RefundedEvent 退款事件
type RefundedEvent struct {
	BaseEvent
	RefundedAmount float64
}
