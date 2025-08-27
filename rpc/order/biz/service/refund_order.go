package service

import (
	"context"
	base "deer/kitex_gen/deer/base"
	order "deer/kitex_gen/deer/order"
)

type RefundOrderService struct {
	ctx context.Context
} // NewRefundOrderService new RefundOrderService
func NewRefundOrderService(ctx context.Context) *RefundOrderService {
	return &RefundOrderService{ctx: ctx}
}

// Run create note info
func (s *RefundOrderService) Run(req *order.RefundOrderReq) (resp *base.BaseResp, err error) {
	// Finish your business logic.

	return
}
