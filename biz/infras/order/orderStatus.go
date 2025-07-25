package order

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

// CanTransitionTo 检查订单状态是否可以转换为目标状态
func (s OrderStatus) CanTransitionTo(target OrderStatus) bool {

	transition := map[OrderStatus][]OrderStatus{
		OrderCreated:   {OrderPaid, OrderCancelled},
		OrderPaid:      {OrderShipped, OrderCancelled, OrderRefunded},
		OrderShipped:   {OrderCompleted, OrderRefunded},
		OrderCancelled: {},
		OrderRefunded:  {},
		OrderCompleted: {},
	}
	for _, allowed := range transition[s] {
		if allowed == target {
			return true
		}
	}

	return true
}
