package common

type OrderStatus string

const (
	Created   OrderStatus = "created"   //创建
	Paying    OrderStatus = "paying"    //支付中
	Paid      OrderStatus = "paid"      //支付完成
	Shipped   OrderStatus = "shipped"   //发货
	Cancelled OrderStatus = "cancelled" //取消
	Completed OrderStatus = "completed" //完成
	Refunded  OrderStatus = "refunded"  //退款

)

// Transitions 定义状态转换规则
var Transitions = map[OrderStatus][]OrderStatus{
	Created:   {Paying, Cancelled},
	Paying:    {Paying, Paid, Cancelled, Refunded},
	Paid:      {Cancelled, Shipped, Refunded},
	Shipped:   {Completed, Refunded},
	Cancelled: {Created},
}

func (s OrderStatus) String() string {
	switch s {
	case Created:
		return "交易已创建，等待用户支付"
	case Paying:
		return "订单支付中"
	case Paid:
		return "订单已支付"
	case Shipped:
		return "发货中"
	case Cancelled:
		return "交易已关闭"
	case Completed:
		return "订单已完成"
	case Refunded:
		return "退款已成功处理"
	default:
		return "状态异常"
	}
}

//1. created "已创建", "交易已创建，等待用户支付"
//2. paying "支付中", "订单支付中"
//3. paid "支付完成", "订单已支付"
//├─ "支付成功", "交易支付成功"
//└─ "支付失败", "交易支付失败"
//4. shipped → "发货"
//├─ "发货成功"
//└─ "发货失败"
//5. cancelled → "已关闭", "交易已关闭"
//6. completed → "完成"
//7. refunded → "退款" "退款处理中", "退款申请已提交，等待处理"
//├─ "退款成功", "退款已成功处理"
//└─ "退款失败", "退款处理失败"
