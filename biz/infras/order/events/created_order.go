package events

import (
	"github.com/google/uuid"
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/status"
	"time"
)

// CreatedOrderEvent 创建订单事件
type CreatedOrderEvent struct {
	common.EventBase
	TotalAmount float64
	Items       []common.Item
	MemberId    int64
	CreatedId   int64
}

func (e *CreatedOrderEvent) GetId() string {
	return uuid.New().String()
}

func (e *CreatedOrderEvent) GetType() string { return string(status.Created) }

func NewCreatedOrderEvent(orderID int64, items []common.Item, amount float64, MemberId int64, userID int64) *CreatedOrderEvent {
	return &CreatedOrderEvent{
		EventBase: common.EventBase{
			EventID:     uuid.New().String(),
			AggregateID: orderID,
			Timestamp:   time.Now(),

			EventType:     "created",
			AggregateType: "order",
			Version:       1,
		},
		TotalAmount: amount,
		Items:       items,
		MemberId:    MemberId,
		CreatedId:   userID,
	}
}
