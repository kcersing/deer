package service

import (
	"context"
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

	return
}
