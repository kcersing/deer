package service

import (
	"context"
	base "gen/kitex_gen/base"
	order "gen/kitex_gen/order"
)

type DeleteOrderService struct {
	ctx context.Context
} // NewDeleteOrderService new DeleteOrderService
func NewDeleteOrderService(ctx context.Context) *DeleteOrderService {
	return &DeleteOrderService{ctx: ctx}
}

// Run create note info
func (s *DeleteOrderService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
