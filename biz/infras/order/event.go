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
	GetVersion() int64
}

// BaseEvent 基础事件结构
type BaseEvent struct {
	EventID     string
	AggregateID int64
	Version     int64
	Type        string
	Timestamp   time.Time
}

func (e *BaseEvent) GetID() string           { return e.EventID }
func (e *BaseEvent) GetType() string         { return e.Type }
func (e *BaseEvent) GetAggregateID() int64   { return e.AggregateID }
func (e *BaseEvent) GetVersion() int64       { return e.Version }
func (e *BaseEvent) GetTimestamp() time.Time { return e.Timestamp }

// CreatedEvent 订单创建事件
type CreatedEvent struct {
	BaseEvent
	OrderSn     string
	Items       []Item
	TotalAmount float64
	CreatedId   int64
}

// 创建事件工厂函数
func NewOrderCreatedEvent(orderID int64, sn string, items []Item, amount float64, userID int64) *CreatedEvent {

	return &CreatedEvent{
		BaseEvent: BaseEvent{
			EventID:     uuid.New().String(),
			AggregateID: orderID,
			Timestamp:   time.Now(),
		},
		OrderSn:     sn,
		Items:       items,
		TotalAmount: amount,
		CreatedId:   userID,
	}
}

func (e *CreatedEvent) GetType() string { return "created" }

// PaidEvent 订单支付事件
type PaidEvent struct {
	BaseEvent
	PayAmount float64
	PayMethod string
}

func (e *PaidEvent) GetType() string { return "paid" }

// CancelledEvent 订单取消事件
type CancelledEvent struct {
	BaseEvent
	Reason string
}

func (e *CancelledEvent) GetType() string { return "cancelled" }

type ShippedEvent struct {
	BaseEvent
}

func (e *ShippedEvent) GetType() string { return "shipped" }

type CompletedEvent struct {
	BaseEvent
}

func (e *CompletedEvent) GetType() string { return "completed" }

type RefundedEvent struct {
	BaseEvent
	RefundedAmount float64
}

func (e *RefundedEvent) GetType() string { return "refunded" }
