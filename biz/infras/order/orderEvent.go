package order

import (
	"github.com/google/uuid"
	"time"
)

// Event 领域事件接口
type Event interface {
	GetID() string
	GetType() string
	GetAggregateID() int64
	GetTimestamp() time.Time
	GetVersion() int
}

// BaseEvent 基础事件结构
type BaseEvent struct {
	EventID     string
	AggregateID int64
	Timestamp   time.Time
	Version     int
}

func (e *BaseEvent) GetID() string           { return e.EventID }
func (e *BaseEvent) GetAggregateID() int64   { return e.AggregateID }
func (e *BaseEvent) GetTimestamp() time.Time { return e.Timestamp }
func (e *BaseEvent) GetVersion() int         { return e.Version }

// OrderCreatedEvent 订单创建事件
type OrderCreatedEvent struct {
	BaseEvent
	OrderSn     string
	Items       []OrderItem
	TotalAmount float64
	CreatedBy   int64
}

func (e *OrderCreatedEvent) GetType() string { return "OrderCreated" }

// OrderPaidEvent 订单支付事件
type OrderPaidEvent struct {
	BaseEvent
	PayAmount float64
	PayMethod string
	PaidAt    time.Time
}

func (e *OrderPaidEvent) GetType() string { return "OrderPaid" }

// OrderCancelledEvent 订单取消事件
type OrderCancelledEvent struct {
	BaseEvent
	Reason      string
	CancelledAt time.Time
}

func (e *OrderCancelledEvent) GetType() string { return "OrderCancelled" }

// 创建事件工厂函数
func NewOrderCreatedEvent(orderID int64, sn string, items []OrderItem, amount float64, userID int64) *OrderCreatedEvent {
	return &OrderCreatedEvent{
		BaseEvent: BaseEvent{
			EventID:     uuid.New().String(),
			AggregateID: orderID,
			Timestamp:   time.Now(),
		},
		OrderSn:     sn,
		Items:       items,
		TotalAmount: amount,
		CreatedBy:   userID,
	}
}

type OrderShippedEvent struct {
	BaseEvent
}

func (e *OrderShippedEvent) GetType() string { return "OrderShipped" }

type OrderCompletedEvent struct {
	BaseEvent
	CompletionAt time.Time
}

func (e *OrderCompletedEvent) GetType() string { return "OrderCompleted" }

type OrderRefundedEvent struct {
	BaseEvent
	RefundAt       time.Time
	RefundedAmount float64
}

func (e *OrderRefundedEvent) GetType() string { return "OrderRefunded" }
