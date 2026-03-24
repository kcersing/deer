package service

import (
	"context"
	base "gen/kitex_gen/base"
	"order/biz/infras/repo"
)

type DeleteOrderService struct {
	ctx context.Context
}

// NewDeleteOrderService new DeleteOrderService
func NewDeleteOrderService(ctx context.Context) *DeleteOrderService {
	return &DeleteOrderService{ctx: ctx}
}

// Run create note info
func (s *DeleteOrderService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	node, err := repo.NewOrderRepo().FindById(s.ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if err = node.Delete(req); err != nil {
		return nil, err
	}

	if err = repo.NewOrderRepo().Save(s.ctx, node); err != nil {
		return nil, err
	}
	return
}
