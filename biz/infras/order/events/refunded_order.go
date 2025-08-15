package events

import (
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/status"
)

// RefundedOrderEvent 退款事件
type RefundedOrderEvent struct {
	common.EventBase
	RefundedAmount float64
	CreatedId      int64
}

func (e *RefundedOrderEvent) GetType() string { return string(status.Refunded) }
