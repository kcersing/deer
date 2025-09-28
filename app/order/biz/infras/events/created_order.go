package events

import (
	"deer/app/order/biz/infras/common"
	"deer/kitex_gen/deer/order"
	"github.com/google/uuid"
	"time"
)

// CreatedOrderEvent 创建订单事件
type CreatedOrderEvent struct {
	common.EventBase
	Sn          string        `json:"Sn,omitempty"`
	TotalAmount float64       `json:"TotalAmount,omitempty"`
	Items       []*order.Item `json:"Items,omitempty"`
	MemberId    int64         `json:"MemberId,omitempty"`
	CreatedId   int64         `json:"CreatedId,omitempty"`
}

func (e *CreatedOrderEvent) GetType() string { return string(common.Created) }

func NewCreatedOrderEvent(sn string, items []*order.Item, amount float64, MemberId int64, userID int64) *CreatedOrderEvent {
	return &CreatedOrderEvent{
		EventBase: common.EventBase{
			EventID:   uuid.New().String(),
			Timestamp: time.Now(),

			EventType:     string(common.Created),
			AggregateType: "order",
			Version:       1,
		},
		Sn:          sn,
		TotalAmount: amount,
		Items:       items,
		MemberId:    MemberId,
		CreatedId:   userID,
	}
}
