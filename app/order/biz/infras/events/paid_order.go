package events

import (
	"common/eventbus"
	"context"
	"order/biz/infras/common"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
)

// PaidOrderEvent 支付订单事件
type PaidOrderEvent struct {
	common.EventBase
}

func (e *PaidOrderEvent) GetType() string { return string(common.Paid) }
func NewPaidOrderEvent(aggregateID int64) *PaidOrderEvent {
	return &PaidOrderEvent{
		EventBase: common.EventBase{
			EventID:       uuid.New().String(),
			AggregateID:   aggregateID,
			Timestamp:     time.Now(),
			EventType:     string(common.Paid),
			AggregateType: "order",
			Version:       1,
		},
	}

}

func HandleOrderPaid(ctx context.Context, req *PaidOrderEvent, event eventbus.Event) error {
	klog.Infof("[Handler] 处理支付完成事件: EventID=%s", event.Id)

	return nil
}
