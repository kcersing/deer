package service

import (
	"context"
	order "order/kitex_gen/order"
)

type GetOrderListService struct {
	ctx context.Context
} // NewGetOrderListService new GetOrderListService
func NewGetOrderListService(ctx context.Context) *GetOrderListService {
	return &GetOrderListService{ctx: ctx}
}

// Run create note info
func (s *GetOrderListService) Run(req *order.GetOrderListReq) (resp *order.GetOrderListResp, err error) {
	// Finish your business logic.

	return
}
