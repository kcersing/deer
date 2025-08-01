package order

import (
	"fmt"
)

// OrderStatus 订单状态
type OrderStatus string

const (
	OrderCreated   OrderStatus = "created"        //创建
	OrderPaid      OrderStatus = "paid"           //支付
	OrderShipped   OrderStatus = "shipped"        //发货
	OrderCancelled OrderStatus = "cancelled"      //取消
	OrderRefunded  OrderStatus = "refunded"       //退款
	OrderCompleted OrderStatus = "OrderCompleted" //完成
)

// OrderStateMachine 订单状态机
type OrderStateMachine struct {
	order *Order
}

func NewOrderStateMachine(order *Order) *OrderStateMachine {
	return &OrderStateMachine{order: order}
}

// 定义状态转换规则
var transitions = map[OrderStatus][]OrderStatus{
	OrderCreated: {OrderPaid, OrderCancelled},
	OrderPaid:    {OrderShipped, OrderRefunded, OrderCancelled},
	OrderShipped: {OrderCompleted, OrderRefunded},
}

// Transition 执行状态转换
func (m *OrderStateMachine) Transition(target OrderStatus, event Event) error {
	m.order.mu.Lock()
	current := m.order.Status
	m.order.mu.Unlock()

	if current == target {
		return nil // 状态未变更
	}

	// 检查是否允许转换
	allowed := false
	for _, status := range transitions[current] {
		if status == target {
			allowed = true
			break
		}
	}
	if !allowed {
		return fmt.Errorf("状态转换不允许: %s -> %s", current, target)
	}

	// 执行转换
	m.order.mu.Lock()
	m.order.Status = target
	m.order.Version++
	m.order.Events = append(m.order.Events, event)
	m.order.mu.Unlock()

	return nil
}
