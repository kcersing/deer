package service

import (
	"context"
	base "gen/kitex_gen/base"
	order "gen/kitex_gen/order"
)

type CancelledOrderService struct {
	ctx context.Context
}

// NewCancelledOrderService new CancelledOrderService
func NewCancelledOrderService(ctx context.Context) *CancelledOrderService {
	return &CancelledOrderService{ctx: ctx}
}

// Run create note info
func (s *CancelledOrderService) Run(req *order.CancelledOrderReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
