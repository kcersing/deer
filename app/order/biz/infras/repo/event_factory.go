package repo

import (
	"fmt"
	"order/biz/infras/common"
	"order/biz/infras/events"
)

// eventFactory 是一个函数，用于创建一个新的事件实例。
type eventFactory func() common.Event

// eventRegistry 是一个事件类型到其工厂函数的映射。
var eventRegistry = make(map[string]eventFactory)

func init() {
	// 在这里注册所有已知的事件类型。
	// 当有新事件时，只需在此处添加一行即可。
	registerEvent(common.Created, func() common.Event { return &events.CreatedOrderEvent{} })
	registerEvent(common.Paid, func() common.Event { return &events.PaidOrderEvent{} })
	registerEvent(common.Shipped, func() common.Event { return &events.ShippedOrderEvent{} })
	registerEvent(common.Cancelled, func() common.Event { return &events.CancelledOrderEvent{} })
	registerEvent(common.Refunded, func() common.Event { return &events.RefundedOrderEvent{} })
	registerEvent(common.Completed, func() common.Event { return &events.CompletedOrderEvent{} })
}

// registerEvent 向注册表中添加一个新的事件工厂。
func registerEvent(eventType common.OrderStatus, factory eventFactory) {
	eventRegistry[string(eventType)] = factory
}

// NewEventByType 根据事件类型字符串创建一个新的事件实例。
// 如果事件类型未注册，则返回错误。
func NewEventByType(eventType string) (common.Event, error) {
	factory, ok := eventRegistry[eventType]
	if !ok {
		return nil, fmt.Errorf("unsupported event type: %s", eventType)
	}
	return factory(), nil
}
