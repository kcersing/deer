package events

import (
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/status"
)

// ShippedOrderEvent 发货事件
type ShippedOrderEvent struct {
	common.EventBase
	CreatedId int64
}

func (e *ShippedOrderEvent) GetType() string { return string(status.Shipped) }
