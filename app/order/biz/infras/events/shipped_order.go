package events

import (
	"common/eventbus"
	"context"
	"gen/kitex_gen/base"
	"gen/kitex_gen/member"
	"order/biz/infras/common"
	"order/rpc/client"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
)

// ShippedOrderEvent 发货事件
type ShippedOrderEvent struct {
	common.EventBase

	MemberId int64

	OrderId int64

	UserId int64

	Items []*base.OrderItem
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
	}
}

func HandleOrderShipped(ctx context.Context, req *ShippedOrderEvent, event eventbus.Event) error {
	klog.Infof("[Handler] 处理发货事件: AggregateID=%s, EventID=%s", req.AggregateID, event)
	//通知生产会员产品
	klog.Infof("[Handler] 处理发货事件: ShippedOrderEvent=%s", req)
	product, err := client.MemberClient.CreateProduct(ctx, &member.CreateProductReq{
		MemberId: req.MemberId,
		OrderId:  req.OrderId,
		UserId:   req.UserId,
		Items:    req.Items,
	})
	if err != nil {
		return err
	}
	klog.Infof("[Handler] 处理发货事件: 返回数据=%s", product)
	//通知生产会员产品

	return nil
}
