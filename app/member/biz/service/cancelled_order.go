package service

import (
	"context"
	base "member/kitex_gen/base"
	order "member/kitex_gen/order"
)

type CancelledOrderService struct {
	ctx context.Context
} // NewCancelledOrderService new CancelledOrderService
func NewCancelledOrderService(ctx context.Context) *CancelledOrderService {
	return &CancelledOrderService{ctx: ctx}
}

// Run create note info
func (s *CancelledOrderService) Run(req *order.CreateOrderReq) (resp *base.BaseResp, err error) {
	// Finish your business logic.

	return
}
