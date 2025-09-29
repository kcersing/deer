package service

import (
	"context"
	base "gen/kitex_gen/base"
)

type DeleteOrderService struct {
	ctx context.Context
} // NewDeleteOrderService new DeleteOrderService
func NewDeleteOrderService(ctx context.Context) *DeleteOrderService {
	return &DeleteOrderService{ctx: ctx}
}

// Run create note info
func (s *DeleteOrderService) Run(req *base.IDReq) (resp *base.BaseResp, err error) {
	// Finish your business logic.

	return
}
