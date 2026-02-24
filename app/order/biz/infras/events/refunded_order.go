package events

import (
	"order/biz/infras/common"
	"time"

	"github.com/google/uuid"
)

// RefundedOrderEvent 退款事件
type RefundedOrderEvent struct {
	common.EventBase
	Reason         string
	RefundedAmount int64
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
