package order

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"sync"
)

// EventHandler 事件处理器接口
type EventHandler interface {
	Handle(event Event) error
}

// EventDispatcher
type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

// NewEventDispatcher
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

// RegisterHandler 注册事件处理器
func (d *EventDispatcher) RegisterHandler(eventType string, handler EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[eventType] = append(d.handlers[eventType], handler)
}

// Dispatch 分发事件
func (d *EventDispatcher) Dispatch(event Event) {
	d.mu.RLock() // 读取锁
	defer d.mu.RUnlock()

	eventType := event.SetEventType()
	if handlers, ok := d.handlers[eventType]; ok {
		var wg sync.WaitGroup
		for _, handler := range handlers {
			wg.Add(1)
			go func(h EventHandler) {
				defer wg.Done()

				if err := h.Handle(event); err != nil {
					// 处理错误
					hlog.Errorf("事件处理失败：%v", err)
				}

			}(handler)
		}

	}

}

// InventoryService 库存服务接口
type InventoryService interface {
	Reserve(sn int64, items []OrderItem) error //预留库存
	Release(sn int64) error                    //扣除库存
	RestoreForRefund(sn int64) error           //恢复库存
}

// InventoryHandler 库存处理器
type InventoryHandler struct {
	InventoryService InventoryService
}

func (h *InventoryHandler) Handle(event Event) error {
	if h.InventoryService == nil {
		return errors.New("库存服务未初始化")
	}
	switch e := event.(type) {
	case *OrderCreatedEvent:
		// 处理订单创建事件
		return h.InventoryService.Reserve(e.Id, e.Items)
	case *OrderCancelledEvent:
		// 处理订单支付事件
		return h.InventoryService.Release(e.Id)
	case *OrderRefundedEvent:
		// 处理订单退款事件
		return h.InventoryService.RestoreForRefund(e.Id)
	default:
		return nil
	}
}

// PayService 支付服务接口
type PayService interface {
	Refund(id int64, amount float64, reason string, createdId int64) error
}
type PayHandler struct {
	PayService PayService
}

func (h *PayHandler) Handle(event Event) error {
	if h.PayService == nil {
		return errors.New("支付服务未初始化")
	}
	switch e := event.(type) {
	case *OrderRefundedEvent:
		return h.PayService.Refund(e.Id, e.RefundedAmount, e.Reason, e.CreatedId)
	default:
		return nil
	}

}

// NotificationServer 通知服务接口
type NotificationServer interface {
	SendOrderCreatedNotification(id int64) error
	SendOrderPaidNotification(id int64) error
	SendOrderShippedNotification(id int64) error
	SendOrderCompletedNotification(id int64) error
	SendOrderCancelledNotification(id int64) error
	SendOrderRefundNotification(id int64) error
}
type NotificationHandler struct {
	NotificationServer NotificationServer
}

func (h *NotificationHandler) Handler(event Event) error {
	if h.NotificationServer == nil {
		return errors.New("通知服务未初始化")
	}
	switch e := event.(type) {
	case *OrderCreatedEvent:
		// 处理订单创建事件
		return h.NotificationServer.SendOrderCreatedNotification(e.Id)
	case *OrderPaidEvent:
		// 处理订单支付事件
		return h.NotificationServer.SendOrderPaidNotification(e.Id)
	case *OrderShippedEvent:
		// 处理订单发货事件
		return h.NotificationServer.SendOrderShippedNotification(e.Id)
	case *OrderCompletedEvent:
		// 处理订单完成事件
		return h.NotificationServer.SendOrderCompletedNotification(e.Id)
	case *OrderCancelledEvent:
		// 处理订单取消事件
		return h.NotificationServer.SendOrderCancelledNotification(e.Id)
	case *OrderRefundedEvent:
		// 处理订单退款事件
		return h.NotificationServer.SendOrderRefundNotification(e.Id)
	default:
		return nil
	}
}
