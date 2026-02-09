package events

import (
	"github.com/google/uuid"
	"order/biz/infras/common"
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
