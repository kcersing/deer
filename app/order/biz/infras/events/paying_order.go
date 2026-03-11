package events

import (
	"common/eventbus"
	"common/pkg/utils"
	"context"
	"gen/kitex_gen/order"
	"order/biz/dal/db"

	"order/biz/infras/common"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
)

// PayingOrderEvent 支付订单事件
type PayingOrderEvent struct {
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

func (e *PayingOrderEvent) GetType() string { return string(common.Paying) }
func NewPayingOrderEvent(req *order.PaymentReq) *PayingOrderEvent {
	return &PayingOrderEvent{
		EventBase: common.EventBase{
			EventID:       uuid.New().String(),
			AggregateID:   req.GetId(),
			Timestamp:     time.Now(),
			EventType:     string(common.Paying),
			AggregateType: "order",
			Version:       1,
		},
		PaySn:     utils.CreateSn(),
		Amount:    req.GetAmount() * 100,
		Method:    req.GetMethod(),
		CreatedId: req.GetUserId(),
		Remission: req.GetRemission() * 100,
		Reason:    req.GetReason(),
		PrepayId:  req.GetPrepayId(),
		PayExtra:  req.GetPayExtra(),
		OrderId:   req.GetId(),
	}

}

func HandleOrderPaying(ctx context.Context, req *PayingOrderEvent, event eventbus.Event) error {
	klog.Infof("[Handler] 处理支付事件: Method=%s, EventID=%s", req.Method, event.Id)
	_, err := db.Client.OrderPay.Create().
		SetOrderID(req.OrderId).
		SetCreatedID(req.CreatedId).
		SetPay(req.Amount * 100).
		SetRemission(req.Remission * 100).
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
