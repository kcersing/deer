package events

import (
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/status"
)

// CancelledOrderEvent 取消订单事件
type CancelledOrderEvent struct {
	common.EventBase
	Reason    string
	CreatedId int64
}

func (e *CancelledOrderEvent) GetType() string { return string(status.Cancelled) }
