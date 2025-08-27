package events

import (
	"deer/biz/infras/common"
	"deer/biz/infras/status"
	"github.com/google/uuid"
	"time"
)

// RefundedOrderEvent 退款事件
type RefundedOrderEvent struct {
	common.EventBase
	RefundedAmount float64
	CreatedId      int64
}

func (e *RefundedOrderEvent) GetType() string { return string(status.Refunded) }
func NewRefundedOrderEvent(AggregateID int64, userID int64) *RefundedOrderEvent {
	return &RefundedOrderEvent{
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
