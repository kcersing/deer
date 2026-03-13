package service

import (
	"context"

	"gen/kitex_gen/member"
)

type ProductListService struct {
	ctx context.Context
}

// NewProductListService new ProductListService
func NewProductListService(ctx context.Context) *ProductListService {
	return &ProductListService{ctx: ctx}
}

// Run create note info
func (s *ProductListService) Run(req *member.ProductListReq) (resp *member.ProductListResp, err error) {
	// Finish your business logic.

	return
}
