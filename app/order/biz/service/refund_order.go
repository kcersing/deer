package service

import (
	"context"
	base "gen/kitex_gen/base"
	order "gen/kitex_gen/order"
	"order/biz/infras/aggregate"
	"order/biz/infras/events"
	"order/biz/infras/repo"
)

type RefundOrderService struct {
	ctx context.Context
}

// NewRefundOrderService new RefundOrderService
func NewRefundOrderService(ctx context.Context) *RefundOrderService {
	return &RefundOrderService{ctx: ctx}
}

// Run create note info
func (s *RefundOrderService) Run(req *order.RefundOrderReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	node := aggregate.NewOrder()
	event := events.NewRefundedOrderEvent(req.GetId(), req.GetCreatedId())
	event.Reason = req.GetReason()
	event.Amount = req.GetAmount()
	err = node.Apply(event)
	if err != nil {
		return nil, err
	}
	err = repo.NewOrderRepo().Save(s.ctx, node)
	if err != nil {
		return nil, err
	}
	
	return

}
