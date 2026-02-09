package events

import (
	"order/biz/infras/common"
	"time"

	"github.com/google/uuid"
)

// PaidOrderEvent 支付订单事件
type PaidOrderEvent struct {
	common.EventBase
	PayedAmount int64
	PayMethod   string
	CreatedId   int64
	Remission   int64
	Reason      string
	PaySn       string
	PrepayId    string
	PayExtra    string
}

func (e *PaidOrderEvent) GetType() string { return string(common.Paid) }
func NewPaidOrderEvent(AggregateID int64, userID int64) *PaidOrderEvent {
	return &PaidOrderEvent{
		EventBase: common.EventBase{
			EventID:     uuid.New().String(),
			AggregateID: AggregateID,
			Timestamp:   time.Now(),

			EventType:     string(common.Paid),
			AggregateType: "order",
			Version:       1,
		},
		CreatedId: userID,
	}
}
