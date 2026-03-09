package service

import (
	"context"
	order "gen/kitex_gen/order"
	"order/biz/infras/events"
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
	event := events.NewPaidOrderEvent(req.GetId())
	event.Amount = req.GetAmount()
	event.Method = req.GetMethod()
	event.PrepayId = req.GetThird()
	err = node.Apply(event)
	if err != nil {
		return nil, err
	}
	err = repo.NewOrderRepo().Save(s.ctx, node)
	if err != nil {
		return nil, err
	}

	err = events.OrderPayEvent(s.ctx, event)
	if err != nil {
		return nil, err
	}

	return &order.OrderResp{
		Data: &node.Order,
	}, nil
}
