package service

import (
	"context"
	base "gen/kitex_gen/base"
	order "gen/kitex_gen/order"
	"order/biz/infras/repo"
)

type RefundOrderService struct {
	ctx context.Context
}

// NewRefundOrderService new RefundOrderService
func NewRefundOrderService(ctx context.Context) *RefundOrderService {
	return &RefundOrderService{ctx: ctx}
}

// Run create note info
func (s *RefundOrderService) Run(req *order.RefundOrderReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	node, err := repo.NewOrderRepo().FindById(s.ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	err = node.Refund(req)
	if err != nil {
		return nil, err
	}
	err = repo.NewOrderRepo().Save(s.ctx, node)
	if err != nil {
		return nil, err
	}
	return
}
