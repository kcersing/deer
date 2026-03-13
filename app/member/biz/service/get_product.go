package service

import (
	"context"
	"gen/kitex_gen/base"
	"gen/kitex_gen/member"
)

type GetProductService struct {
	ctx context.Context
}

// NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *base.IdReq) (resp *member.ProductResp, err error) {
	// Finish your business logic.

	return
}
