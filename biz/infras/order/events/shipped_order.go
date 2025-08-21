package events

import (
	"github.com/google/uuid"
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/status"
	"time"
)

// ShippedOrderEvent 发货事件
type ShippedOrderEvent struct {
	common.EventBase
	CreatedId int64
}

func (e *ShippedOrderEvent) GetType() string { return string(status.Shipped) }

func NewShippedOrderEvent(AggregateID int64, userID int64) *ShippedOrderEvent {
	return &ShippedOrderEvent{
		EventBase: common.EventBase{
			EventID:     uuid.New().String(),
			AggregateID: AggregateID,
			Timestamp:   time.Now(),

			EventType:     string(status.Shipped),
			AggregateType: "order",
			Version:       1,
		},
		CreatedId: userID,
	}
}
