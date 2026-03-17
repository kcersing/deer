package events

import (
	"common/eventbus"
	"context"
	"order/biz/infras/common"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
)

// ShippedOrderEvent 发货事件
type ShippedOrderEvent struct {
	common.EventBase
	CreatedId int64
}

func (e *ShippedOrderEvent) GetType() string { return string(common.Shipped) }

func NewShippedOrderEvent(aggregateID int64) *ShippedOrderEvent {
	return &ShippedOrderEvent{
		EventBase: common.EventBase{
			EventID:     uuid.New().String(),
			AggregateID: aggregateID,
			Timestamp:   time.Now(),

			EventType:     string(common.Shipped),
			AggregateType: "order",
			Version:       1,
		},
		//CreatedId: req.userID,
	}
}

func HandleOrderShipped(ctx context.Context, req *ShippedOrderEvent, event eventbus.Event) error {
	klog.Infof("[Handler] 处理发货事件: AggregateID=%s, EventID=%s", req.AggregateID, event.Id)
	//通知生产会员产品

	return nil
}
