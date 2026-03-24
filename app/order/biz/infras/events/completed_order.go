package events

import (
	"gen/kitex_gen/order"
	"order/biz/infras/common"
	"time"

	"github.com/google/uuid"
)

// CompletedOrderEvent 完成订单事件
type CompletedOrderEvent struct {
	common.EventBase
	CreatedId int64
}

func (e *CompletedOrderEvent) GetType() string { return string(common.Completed) }

func NewCompletedOrderEvent(req *order.CompletedOrderReq) *CompletedOrderEvent {
	return &CompletedOrderEvent{
		EventBase: common.EventBase{
			EventID:     uuid.New().String(),
			AggregateID: req.GetId(),
			Timestamp:   time.Now(),

			EventType:     string(common.Completed),
			AggregateType: "order",
			Version:       1,
		},
		CreatedId: req.GetCreatedId(),
	}
}
