package service

import (
	"context"
	base "gen/kitex_gen/base"
	order "gen/kitex_gen/order"
	"order/biz/infras/repo"
)

type CancelledOrderService struct {
	ctx context.Context
}

// NewCancelledOrderService new CancelledOrderService
func NewCancelledOrderService(ctx context.Context) *CancelledOrderService {
	return &CancelledOrderService{ctx: ctx}
}

// Run create note info 取消订单
func (s *CancelledOrderService) Run(req *order.CancelledOrderReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	node, err := repo.NewOrderRepo().FindById(s.ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if err = node.Cancel(req); err != nil {
		return nil, err
	}
	if err = repo.NewOrderRepo().Save(s.ctx, node); err != nil {
		return nil, err
	}
	return
}
