package events

import (
	"common/eventbus"
	"common/pkg/utils"
	"context"
	"order/biz/dal/db"

	"order/biz/infras/common"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
)

// PaidOrderEvent 支付订单事件
type PaidOrderEvent struct {
	common.EventBase
	Amount    int64
	Method    string
	CreatedId int64
	Remission int64
	Reason    string
	PaySn     string
	PrepayId  string
	PayExtra  string
	OrderId   int64
}

func (e *PaidOrderEvent) GetType() string { return string(common.Paid) }
func NewPaidOrderEvent(AggregateID int64) *PaidOrderEvent {
	return &PaidOrderEvent{
		EventBase: common.EventBase{
			EventID:       uuid.New().String(),
			AggregateID:   AggregateID,
			Timestamp:     time.Now(),
			EventType:     string(common.Paid),
			AggregateType: "order",
			Version:       1,
		},
		PaySn: utils.CreateSn(),
	}

}

func HandleOrderPay(ctx context.Context, req *PaidOrderEvent, event eventbus.Event) error {
	klog.Infof("[Handler] 处理支付事件: Method=%s, EventID=%s", req.Method, event.Id)
	_, err := db.Client.OrderPay.Create().
		SetOrderID(req.OrderId).
		SetCreatedID(req.CreatedId).
		SetPay(req.Amount).
		SetNote(req.Reason).
		SetPayAt(time.Now()).
		SetPayWay(req.Method).
		SetPaySn(req.PaySn).
		SetPrepayID(req.PrepayId).
		SetPayExtra(req.PayExtra).
		Save(ctx)
	if err != nil {
		klog.Warnf("[Handler] 保存订单支付信息失败: %v", err)
		return err
	}
	return nil
}
