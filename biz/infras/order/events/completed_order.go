package events

import (
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/status"
)

// CompletedOrderEvent 完成订单事件
type CompletedOrderEvent struct {
	common.EventBase
	CreatedId int64
}

func (e *CompletedOrderEvent) GetType() string { return string(status.Completed) }
