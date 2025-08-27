package events

import (
	"deer/biz/infras/common"
	"deer/biz/infras/status"
	"github.com/google/uuid"
	"time"
)

// CancelledOrderEvent 取消订单事件
type CancelledOrderEvent struct {
	common.EventBase
	Reason    string
	CreatedId int64
}

func (e *CancelledOrderEvent) GetType() string { return string(status.Cancelled) }

func NewCancelledOrderEvent(AggregateID int64, reason string, userID int64) *CancelledOrderEvent {
	return &CancelledOrderEvent{
		EventBase: common.EventBase{
			EventID:     uuid.New().String(),
			AggregateID: AggregateID,
			Timestamp:   time.Now(),

			EventType:     string(status.Cancelled),
			AggregateType: "order",
			Version:       1,
		},
		Reason:    reason,
		CreatedId: userID,
	}
}
