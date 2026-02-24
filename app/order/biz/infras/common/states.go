package common

type OrderStatus string

const (
	Created   OrderStatus = "created"   //创建
	Paid      OrderStatus = "paid"      //支付
	Shipped   OrderStatus = "shipped"   //发货
	Cancelled OrderStatus = "cancelled" //取消
	Completed OrderStatus = "completed" //完成
	Refunded  OrderStatus = "refunded"  //退款

)

// Transitions 定义状态转换规则
var Transitions = map[OrderStatus][]OrderStatus{
	Created: {Paid, Cancelled},
	Paid:    {Shipped, Refunded, Cancelled},
	Shipped: {Completed, Refunded},
}

//1. created "已创建", "交易已创建，等待用户支付"
//2. paid "支付", "用户已发起支付，等待支付结果"
//├─ "支付成功", "交易支付成功"
//└─ "支付失败", "交易支付失败"
//3. shipped → "发货"
//├─ "发货成功"
//└─ "发货失败"
//4. cancelled → "已关闭", "交易已关闭"
//5. completed → "完成"
//6. refunded → "退款" "退款处理中", "退款申请已提交，等待处理"
//├─ "退款成功", "退款已成功处理"
//└─ "退款失败", "退款处理失败"
