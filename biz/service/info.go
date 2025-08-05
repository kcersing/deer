package service

import (
	"context"
	order2 "kcers-order/biz/infras/order"
	order "kcers-order/kitex_gen/order"
)

type InfoService struct {
	ctx context.Context
} // NewInfoService new InfoService
func NewInfoService(ctx context.Context) *InfoService {
	return &InfoService{ctx: ctx}
}

// Run create note info
func (s *InfoService) Run(req *order.Req) (resp *order.Resp, err error) {
	// Finish your business logic.

	order := order2.NewOrder(1, "1111", nil, 100)

	createdEvent := order2.NewOrderCreatedEvent(1, "1111", nil, 100, 10000)

	order.AddEvent(createdEvent)

	return
}
