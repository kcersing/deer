package events

import (
	"gen/kitex_gen/order"
	"order/biz/infras/common"
	"time"

	"github.com/google/uuid"
)

// CancelledOrderEvent 取消订单事件
type CancelledOrderEvent struct {
	common.EventBase
	Reason    string
	CreatedId int64
}

func (e *CancelledOrderEvent) GetType() string { return string(common.Cancelled) }

func NewCancelledOrderEvent(req *order.CancelledOrderReq) *CancelledOrderEvent {
	return &CancelledOrderEvent{
		EventBase: common.EventBase{
			EventID:     uuid.New().String(),
			AggregateID: req.GetId(),
			Timestamp:   time.Now(),

			EventType:     string(common.Cancelled),
			AggregateType: "order",
			Version:       1,
		},
		Reason:    req.GetReason(),
		CreatedId: req.GetCreatedId(),
	}
}
