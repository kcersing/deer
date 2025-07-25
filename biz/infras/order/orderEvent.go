package order

import "time"

// OrderCreatedEvent 订单创建事件
type OrderCreatedEvent struct {
	Id          int64
	EventId     string
	Items       []OrderItem
	TotalAmount float64
	CreatedAt   time.Time
	CreatedId   int64
}

func (e *OrderCreatedEvent) GetEventId() string {
	return e.EventId
}
func (e *OrderCreatedEvent) GetId() int64 {
	return e.Id
}
func (e *OrderCreatedEvent) GetCreatedAt() time.Time {
	return e.CreatedAt
}
func (e *OrderCreatedEvent) GetEventType() string {
	return "OrderCreated"
}

// OrderPaidEvent 订单支付事件
type OrderPaidEvent struct {
	EventId   string
	Id        int64
	PaidAt    time.Time
	PayMethod string
	PayAmount float64
	CreatedAt time.Time
	CreatedId int64
}

func (e *OrderPaidEvent) GetEventId() string {
	return e.EventId
}
func (e *OrderPaidEvent) GetId() int64 {
	return e.Id
}
func (e *OrderPaidEvent) GetCreatedAt() time.Time {
	return e.CreatedAt
}
func (e *OrderPaidEvent) GetEventType() string {
	return "OrderPaid"
}

// OrderShippedEvent 订单发货事件
type OrderShippedEvent struct {
	EventId   string
	Id        int64
	ShippedAt time.Time
	CreatedAt time.Time
	CreatedId int64
}

func (e *OrderShippedEvent) GetEventId() string {
	return e.EventId
}
func (e *OrderShippedEvent) GetId() int64 {
	return e.Id
}
func (e *OrderShippedEvent) GetShippedAt() time.Time {
	return e.ShippedAt
}
func (e *OrderShippedEvent) GetCreatedAt() time.Time {
	return e.CreatedAt
}
func (e *OrderShippedEvent) GetEventType() string {
	return "OrderShipped"
}

// OrderCompletedEvent 订单完成事件
type OrderCompletedEvent struct {
	EventId     string
	Id          int64
	CompletedAt time.Time
	CreatedAt   time.Time
	CreatedId   int64
}

func (e *OrderCompletedEvent) GetEventId() string {
	return e.EventId
}
func (e *OrderCompletedEvent) GetId() int64 {
	return e.Id
}
func (e *OrderCompletedEvent) GetCompletedAt() time.Time {
	return e.CompletedAt
}
func (e *OrderCompletedEvent) GetCreatedAt() time.Time {
	return e.CreatedAt
}
func (e *OrderCompletedEvent) GetEventType() string {
	return "OrderCompleted"
}

// OrderCancelledEvent 订单取消事件
type OrderCancelledEvent struct {
	EventId     string
	Id          int64
	CancelledAt time.Time
	CreatedAt   time.Time
	CreatedId   int64
}

func (e *OrderCancelledEvent) GetEventId() string {
	return e.EventId
}
func (e *OrderCancelledEvent) GetId() int64 {
	return e.Id
}
func (e *OrderCancelledEvent) GetCancelledAt() time.Time {
	return e.CancelledAt
}
func (e *OrderCancelledEvent) GetCreatedAt() time.Time {
	return e.CreatedAt
}
func (e *OrderCancelledEvent) GetEventType() string {
	return "OrderCancelled"
}

// OrderRefundedEvent 订单退款事件
type OrderRefundedEvent struct {
	EventId        string
	Id             int64
	RefundedAt     time.Time
	CreatedAt      time.Time
	CreatedId      int64
	RefundedAmount float64
	Reason         string //退款原因
}

func (e *OrderRefundedEvent) GetEventId() string {
	return e.EventId
}
func (e *OrderRefundedEvent) GetId() int64 {
	return e.Id
}
func (e *OrderRefundedEvent) GetRefundedAt() time.Time {
	return e.RefundedAt
}
func (e *OrderRefundedEvent) GetCreatedAt() time.Time {
	return e.CreatedAt
}
func (e *OrderRefundedEvent) GetEventType() string {
	return "OrderRefunded"
}
