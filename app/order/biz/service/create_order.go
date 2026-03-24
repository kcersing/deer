package service

import (
	"common/eventbus"
	"context"
	order "gen/kitex_gen/order"
	"order/biz/infras"
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
	return nil, err
}
