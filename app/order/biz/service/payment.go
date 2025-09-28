package service

import (
	"context"
	order "order/kitex_gen/order"
)

type PaymentService struct {
	ctx context.Context
} // NewPaymentService new PaymentService
func NewPaymentService(ctx context.Context) *PaymentService {
	return &PaymentService{ctx: ctx}
}

// Run create note info
func (s *PaymentService) Run(req *order.PaymentReq) (resp *order.OrderResp, err error) {
	// Finish your business logic.

	return
}
