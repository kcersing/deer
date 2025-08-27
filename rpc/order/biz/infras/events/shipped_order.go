package events

import (
	"deer/rpc/order/biz/infras/common"
	"github.com/google/uuid"
	"time"
)

// ShippedOrderEvent 发货事件
type ShippedOrderEvent struct {
	common.EventBase
	CreatedId int64
}

func (e *ShippedOrderEvent) GetType() string { return string(common.Shipped) }

func NewShippedOrderEvent(AggregateID int64, userID int64) *ShippedOrderEvent {
	return &ShippedOrderEvent{
		EventBase: common.EventBase{
			EventID:     uuid.New().String(),
			AggregateID: AggregateID,
			Timestamp:   time.Now(),

			EventType:     string(common.Shipped),
			AggregateType: "order",
			Version:       1,
		},
		CreatedId: userID,
	}
}
