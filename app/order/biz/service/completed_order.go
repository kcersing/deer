package service

import (
	"context"
	base "gen/kitex_gen/base"
	order "gen/kitex_gen/order"
	"order/biz/infras/repo"
)

type CompletedOrderService struct {
	ctx context.Context
}

// NewCompletedOrderService new CompletedOrderService
func NewCompletedOrderService(ctx context.Context) *CompletedOrderService {
	return &CompletedOrderService{ctx: ctx}
}

// Run create note info
func (s *CompletedOrderService) Run(req *order.CompletedOrderReq) (resp *base.NilResponse, err error) {

	node, err := repo.NewOrderRepo().FindById(s.ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if err = node.Completed(req); err != nil {
		return nil, err
	}

	if err = repo.NewOrderRepo().Save(s.ctx, node); err != nil {
		return nil, err
	}
	return
}
