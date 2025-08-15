package events

import (
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/status"
)

// PaidOrderEvent 支付订单事件
type PaidOrderEvent struct {
	common.EventBase
	PayedAmount float64
	PayMethod   string
	CreatedId   int64
}

func (e *PaidOrderEvent) GetType() string { return string(status.Paid) }
