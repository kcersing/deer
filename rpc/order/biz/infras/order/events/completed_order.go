package events

import (
	"deer/biz/infras/common"
	"deer/biz/infras/status"
	"github.com/google/uuid"
	"time"
)

// CompletedOrderEvent 完成订单事件
type CompletedOrderEvent struct {
	common.EventBase
	CreatedId int64
}

func (e *CompletedOrderEvent) GetType() string { return string(status.Completed) }

func NewCompletedOrderEvent(AggregateID int64, userID int64) *CompletedOrderEvent {
	return &CompletedOrderEvent{
		EventBase: common.EventBase{
			EventID:     uuid.New().String(),
			AggregateID: AggregateID,
			Timestamp:   time.Now(),

			EventType:     string(status.Refunded),
			AggregateType: "order",
			Version:       1,
		},
		CreatedId: userID,
	}
}
