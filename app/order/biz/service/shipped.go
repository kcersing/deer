package service

import (
	"context"
	base "gen/kitex_gen/base"
	"order/biz/infras/repo"
)

type ShippedService struct {
	ctx context.Context
}

// NewShippedService new ShippedService
func NewShippedService(ctx context.Context) *ShippedService {
	return &ShippedService{ctx: ctx}
}

// Run create note info
func (s *ShippedService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	node, err := repo.NewOrderRepo().FindById(s.ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	err = node.Shipped()
	if err != nil {
		return nil, err
	}
	err = repo.NewOrderRepo().Save(s.ctx, node)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
