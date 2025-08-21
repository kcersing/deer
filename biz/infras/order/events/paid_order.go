package events

import (
	"github.com/google/uuid"
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/status"
	"time"
)

// PaidOrderEvent 支付订单事件
type PaidOrderEvent struct {
	common.EventBase
	PayedAmount float64
	PayMethod   string
	CreatedId   int64
}

func (e *PaidOrderEvent) GetType() string { return string(status.Paid) }
func NewPaidOrderEvent(AggregateID int64, userID int64) *PaidOrderEvent {
	return &PaidOrderEvent{
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
