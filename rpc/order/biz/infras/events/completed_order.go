package events

import (
	"deer/rpc/order/biz/infras/common"
	"github.com/google/uuid"
	"time"
)

// CompletedOrderEvent 完成订单事件
type CompletedOrderEvent struct {
	common.EventBase
	CreatedId int64
}

func (e *CompletedOrderEvent) GetType() string { return string(common.Completed) }

func NewCompletedOrderEvent(AggregateID int64, userID int64) *CompletedOrderEvent {
	return &CompletedOrderEvent{
		EventBase: common.EventBase{
			EventID:     uuid.New().String(),
			AggregateID: AggregateID,
			Timestamp:   time.Now(),

			EventType:     string(common.Completed),
			AggregateType: "order",
			Version:       1,
		},
		CreatedId: userID,
	}
}
