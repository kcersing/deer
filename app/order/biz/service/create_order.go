package service

import (
	"context"
	base "gen/kitex_gen/base"
	order "gen/kitex_gen/order"
	"order/biz/infras/aggregate"
	"order/biz/infras/events"
	"order/biz/infras/repo"
)

type CreateOrderService struct {
	ctx context.Context
}

// NewCreateOrderService new CreateOrderService
func NewCreateOrderService(ctx context.Context) *CreateOrderService {
	return &CreateOrderService{ctx: ctx}
}

// Run create note info
func (s *CreateOrderService) Run(req *order.CreateOrderReq) (resp *order.OrderResp, err error) {
	// Finish your business logic.
	order := aggregate.NewOrder()

	event := events.NewCreatedOrderEvent(
		"SN-20240101-001",
		[]*base.OrderItem{},
		999.9,
		req.GetMemberId(),
		req.GetUserId(),
	)

	order.Apply(event) // 订单应用事件，改变状态为 "created"
	repo.Save(order)   // 持久化（写事件表、聚合表、快照）
	// 事务提交后 → 发布到 eventbus → 库存预留、发送通知等异步处理
	return
}
