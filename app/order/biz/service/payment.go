package service

import (
	"context"
	order "gen/kitex_gen/order"
)

type PaymentService struct {
	ctx context.Context
}

// NewPaymentService new PaymentService
func NewPaymentService(ctx context.Context) *PaymentService {
	return &PaymentService{ctx: ctx}
}

// Run create note info
func (s *PaymentService) Run(req *order.PaymentReq) (resp *order.OrderResp, err error) {
	// Finish your business logic.
	//// === 第三步：支付订单 ===
	//paidEvent := events.NewPaidOrderEvent(orderId, userId)
	//paidEvent.PayedAmount = 9999
	//paidEvent.PayMethod = "alipay"
	//
	//orderFromDB.Apply(paidEvent)  // 订单状态转为 "paid"，添加支付记录到 OrderPays
	//repo.Save(orderFromDB)        // 持久化新事件（版本 2）
	//// 事务提交后 → 发布支付事件到 eventbus → 触发发货流程
	return
}
