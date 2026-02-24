package service

import (
	"common/pkg/utils"
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
	items := make([]*base.OrderItem, 0, len(req.GetItems()))

	for _, item := range req.GetItems() {
		items = append(items, &base.OrderItem{
			ProductId: item.GetProductId(),
			Quantity:  item.GetQuantity(),
		})
	}
	event := events.NewCreatedOrderEvent(
		utils.CreateCn(),
		items,
		// 金额单位：分
		req.GetTotalAmount()*100,
		req.GetMemberId(),
		req.GetUserId(),
	)

	err = order.Apply(event)
	if err != nil {
		return nil, err
	} // 订单应用事件，改变状态为 "created"
	err = repo.OrderRepoClient.Save(order)
	if err != nil {
		return nil, err
	} // 持久化（写事件表、聚合表、快照）
	// 事务提交后 → 发布到 eventbus → 库存预留、发送通知等异步处理
	return
}
