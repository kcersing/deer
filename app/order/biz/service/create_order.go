package service

import (
	"common/eventbus"
	"common/pkg/utils"
	"context"
	base "gen/kitex_gen/base"
	order "gen/kitex_gen/order"
	"order/biz/infras"
	"order/biz/infras/aggregate"
	"order/biz/infras/events"
	"order/biz/infras/repo"
)

type CreateOrderService struct {
	ctx       context.Context
	publisher *eventbus.EventPublisher
}

// NewCreateOrderService new CreateOrderService
func NewCreateOrderService(ctx context.Context) *CreateOrderService {
	return &CreateOrderService{
		ctx:       ctx,
		publisher: infras.GetManager().Publisher,
	}
}

// Run create note info
func (s *CreateOrderService) Run(req *order.CreateOrderReq) (resp *order.OrderResp, err error) {
	node := aggregate.NewOrder()
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
		req.GetTotalAmount()*100,
		req.GetMemberId(),
		req.GetUserId(),
	)

	err = node.Apply(event)
	if err != nil {
		return nil, err
	}

	// 将 publisher 注入 Repository，由 Repository 负责在事务提交后发布事件
	err = repo.NewOrderRepo(s.publisher).Save(s.ctx, node)
	if err != nil {
		return nil, err
	}

	return
}
