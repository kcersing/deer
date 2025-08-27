package events

import (
	"deer/biz/infras/common"
	"deer/biz/infras/status"
	"github.com/google/uuid"
	"time"
)

// CreatedOrderEvent 创建订单事件
type CreatedOrderEvent struct {
	common.EventBase
	Sn          string
	TotalAmount float64
	Items       []common.Item
	MemberId    int64
	CreatedId   int64
}

func (e *CreatedOrderEvent) GetType() string { return string(status.Created) }

func NewCreatedOrderEvent(sn string, items []common.Item, amount float64, MemberId int64, userID int64) *CreatedOrderEvent {
	return &CreatedOrderEvent{
		EventBase: common.EventBase{
			EventID:   uuid.New().String(),
			Timestamp: time.Now(),

			EventType:     string(status.Created),
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
