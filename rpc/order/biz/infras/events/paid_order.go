package events

import (
	"deer/rpc/order/biz/infras/common"
	"github.com/google/uuid"
	"time"
)

// PaidOrderEvent 支付订单事件
type PaidOrderEvent struct {
	common.EventBase
	PayedAmount float64
	PayMethod   string
	CreatedId   int64
	Remission   float64
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
