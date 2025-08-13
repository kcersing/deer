package service

import (
	"fmt"
	"kcers-order/biz/infras/aggregate"
	"kcers-order/biz/infras/events"
	"kcers-order/biz/infras/states"
)

// StateMachine 订单状态机
type StateMachine struct {
	order *aggregate.Order
}

func NewStateMachine(order *aggregate.Order) *StateMachine {
	return &StateMachine{order: order}
}

// 定义状态转换规则
var transitions = map[states.OrderStatus][]states.OrderStatus{
	states.Created: {states.Paid, states.Cancelled},
	states.Paid:    {states.Shipped, states.Refunded, states.Cancelled},
	states.Shipped: {states.Completed, states.Refunded},
}

// Transition 执行状态转换
func (m *StateMachine) Transition(target states.OrderStatus, event events.Event) error {
	m.order.Mu.Lock()
	current := m.order.Status
	m.order.Mu.Unlock()

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
	m.order.Mu.Lock()
	m.order.Status = target
	m.order.Version++
	m.order.Events = append(m.order.Events, event)
	m.order.Mu.Unlock()

	return nil
}
