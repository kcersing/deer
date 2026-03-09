package events

import (
	"common/eventbus"
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"

	"order/biz/dal/db"
	"order/biz/infras/common"
	"time"
)

// RefundedOrderEvent 退款事件
type RefundedOrderEvent struct {
	common.EventBase
	Reason    string
	Amount    int64
	CreatedId int64
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

func HandleOrderRefunded(ctx context.Context, req *RefundedOrderEvent, event eventbus.Event) error {
	klog.Infof("[Handler] 处理退费事件: Reason=%s, EventID=%s", req.Reason, event.Id)
	_, err := db.Client.OrderRefund.Create().
		SetOrderID(req.GetAggregateID()).
		SetCreatedID(req.CreatedId).
		SetRefundAt(time.Now()).
		SetNature(req.Reason).
		SetAmount(req.Amount).
		Save(ctx)

	if err != nil {
		klog.Warnf("[Handler] 保存订单退费信息失败: %v", err)
		return err
	}
	return nil
}
