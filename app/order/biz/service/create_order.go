package service

import (
	"common/eventbus"
	"context"
	order "gen/kitex_gen/order"
	"order/biz/infras"
	"order/biz/infras/aggregate"
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
	err = node.Create(req)
	if err != nil {
		return nil, err
	}
	err = repo.NewOrderRepo().Save(s.ctx, node)
	if err != nil {
		return nil, err
	}

	return &order.OrderResp{
		Data: &node.Order,
	}, nil
}
