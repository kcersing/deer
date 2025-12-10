package service

import (
	"context"
	base "gen/kitex_gen/base"
	order "gen/kitex_gen/order"
)

type RefundOrderService struct {
	ctx context.Context
} // NewRefundOrderService new RefundOrderService
func NewRefundOrderService(ctx context.Context) *RefundOrderService {
	return &RefundOrderService{ctx: ctx}
}

// Run create note info
func (s *RefundOrderService) Run(req *order.RefundOrderReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
