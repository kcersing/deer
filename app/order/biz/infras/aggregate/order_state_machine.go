package aggregate

import (
	"fmt"
	"order/biz/infras/common"
	"slices"

	"github.com/cloudwego/kitex/pkg/klog"
)

// StateMachine 订单状态机
type StateMachine struct {
	order *Order
}

func NewStateMachine(order *Order) *StateMachine {
	return &StateMachine{order: order}
}

// ValidateTransition 校验状态转换是否允许
func (m *StateMachine) ValidateTransition(target common.OrderStatus) error {

	current := m.order.Status
	if current == string(target) {
		return fmt.Errorf("状态未变更: %s -> %s", current, target)
	}
	// 检查是否允许转换
	allowed := slices.Contains(common.Transitions[common.OrderStatus(current)], target)

	klog.Infof("target", target)
	klog.Infof("current", current)
	klog.Infof("common.OrderStatus(current)", common.OrderStatus(current))
	klog.Infof("common.Transitions[common.OrderStatus(current)]", common.Transitions[common.OrderStatus(current)])
	klog.Infof("结果", allowed)
	if !allowed {
		return fmt.Errorf("状态转换不允许: %s -> %s", current, target)
	}

	return nil
}
