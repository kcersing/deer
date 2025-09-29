package service

import (
	"context"
	order "member/kitex_gen/order"
)

type CreateOrderService struct {
	ctx context.Context
} // NewCreateOrderService new CreateOrderService
func NewCreateOrderService(ctx context.Context) *CreateOrderService {
	return &CreateOrderService{ctx: ctx}
}

// Run create note info
func (s *CreateOrderService) Run(req *order.GetOrderListReq) (resp *order.OrderResp, err error) {
	// Finish your business logic.

	return
}
