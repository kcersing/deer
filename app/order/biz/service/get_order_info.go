package service

import (
	"context"
	base "gen/kitex_gen/base"
	order "gen/kitex_gen/order"
)

type GetOrderInfoService struct {
	ctx context.Context
}

// NewGetOrderInfoService new GetOrderInfoService
func NewGetOrderInfoService(ctx context.Context) *GetOrderInfoService {
	return &GetOrderInfoService{ctx: ctx}
}

// Run create note info
func (s *GetOrderInfoService) Run(req *order.GetOrderInfoReq) (resp *order.OrderResp, err error) {
	// Finish your business logic.

	return
}
