package events

import (
	"common/pkg/utils"
	"gen/kitex_gen/base"
	"gen/kitex_gen/order"
	"order/biz/infras/common"
	"time"

	"github.com/google/uuid"
)

// CreatedOrderEvent 创建订单事件
type CreatedOrderEvent struct {
	common.EventBase
	Sn          string            `json:"Sn,omitempty"`
	TotalAmount int64             `json:"TotalAmount,omitempty"`
	Items       []*base.OrderItem `json:"Items,omitempty"`
	MemberId    int64             `json:"MemberId,omitempty"`
	CreatedId   int64             `json:"CreatedId,omitempty"`
}

func (e *CreatedOrderEvent) GetType() string { return string(common.Created) }

func NewCreatedOrderEvent(req *order.CreateOrderReq) *CreatedOrderEvent {

	items := make([]*base.OrderItem, 0, len(req.GetItems()))
	for _, item := range req.GetItems() {
		items = append(items, &base.OrderItem{
			ProductId: item.GetProductId(),
			Quantity:  item.GetQuantity(),
			Price:     item.GetPrice(),
			Name:      item.GetName(),
		})
	}

	return &CreatedOrderEvent{
		EventBase: common.EventBase{
			EventID:   uuid.New().String(),
			Timestamp: time.Now(),

			EventType:     string(common.Created),
			AggregateType: "order",
			Version:       1,
		},
		Sn:          utils.CreateSn(),
		TotalAmount: req.GetTotalAmount() * 100,
		Items:       items,
		MemberId:    req.GetMemberId(),
		CreatedId:   req.GetUserId(),
	}
}
