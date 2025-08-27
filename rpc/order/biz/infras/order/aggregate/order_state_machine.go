package aggregate

import (
	"deer/biz/infras/common"
	"deer/biz/infras/status"
	"fmt"
)

// StateMachine 订单状态机
type StateMachine struct {
	order *Order
}

func NewStateMachine(order *Order) *StateMachine {
	return &StateMachine{order: order}
}

// 定义状态转换规则
var transitions = map[status.OrderStatus][]status.OrderStatus{
	status.Created: {status.Paid, status.Cancelled},
	status.Paid:    {status.Shipped, status.Refunded, status.Cancelled},
	status.Shipped: {status.Completed, status.Refunded},
}

// Transition 执行状态转换
func (m *StateMachine) Transition(target status.OrderStatus, event common.Event) error {
	m.order.mu.Lock()
	current := m.order.Status
	m.order.mu.Unlock()

	if current == target {
		return nil // 状态未变更
	}

	// 检查是否允许转换
	allowed := false
	for _, statu := range transitions[current] {
		if statu == target {
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
	//m.order.uncommittedEvents = append(m.order.uncommittedEvents, event)
	m.order.mu.Unlock()

	return nil
}
