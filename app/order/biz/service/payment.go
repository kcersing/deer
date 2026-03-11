package service

import (
	"context"
	order "gen/kitex_gen/order"
	"order/biz/infras/repo"
)

type PaymentService struct {
	ctx context.Context
}

// NewPaymentService new PaymentService
func NewPaymentService(ctx context.Context) *PaymentService {
	return &PaymentService{ctx: ctx}
}

// Run create note info
func (s *PaymentService) Run(req *order.PaymentReq) (resp *order.OrderResp, err error) {

	node, err := repo.NewOrderRepo().FindById(s.ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	err = node.Paying(req)

	if err != nil {
		return nil, err
	}
	err = repo.NewOrderRepo().Save(s.ctx, node)
	if err != nil {
		return nil, err
	}

	return &order.OrderResp{
		Data: &node.Order,
	}, nil
}
