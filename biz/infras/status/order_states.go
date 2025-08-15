package status

type OrderStatus string

const (
	Created   OrderStatus = "created"   //创建
	Paid      OrderStatus = "paid"      //支付
	Shipped   OrderStatus = "shipped"   //发货
	Cancelled OrderStatus = "cancelled" //取消
	Refunded  OrderStatus = "refunded"  //退款
	Completed OrderStatus = "completed" //完成
)
