package events

import (
	"github.com/google/uuid"
	"order/biz/infras/common"
	"time"
)

// RefundedOrderEvent 退款事件
type RefundedOrderEvent struct {
	common.EventBase
	Reason         string
	RefundedAmount float64
	CreatedId      int64
}

func (e *RefundedOrderEvent) GetType() string { return string(common.Refunded) }
func NewRefundedOrderEvent(AggregateID int64, userID int64) *RefundedOrderEvent {
	return &RefundedOrderEvent{
		EventBase: common.EventBase{
			EventID:     uuid.New().String(),
			AggregateID: AggregateID,
			Timestamp:   time.Now(),

			EventType:     string(common.Refunded),
			AggregateType: "order",
			Version:       1,
		},
		CreatedId: userID,
	}
}
