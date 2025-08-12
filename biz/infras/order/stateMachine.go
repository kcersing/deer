package order

import (
	"fmt"
)

// Status OrderStatus 订单状态
type Status string

const (
	Created   Status = "created"   //创建
	Paid      Status = "paid"      //支付
	Shipped   Status = "shipped"   //发货
	Cancelled Status = "cancelled" //取消
	Refunded  Status = "refunded"  //退款
	Completed Status = "completed" //完成
)

// StateMachine 订单状态机
type StateMachine struct {
	order *Order
}

func NewStateMachine(order *Order) *StateMachine {
	return &StateMachine{order: order}
}

// 定义状态转换规则
var transitions = map[Status][]Status{
	Created: {Paid, Cancelled},
	Paid:    {Shipped, Refunded, Cancelled},
	Shipped: {Completed, Refunded},
}

// Transition 执行状态转换
func (m *StateMachine) Transition(target Status, event Event) error {
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
