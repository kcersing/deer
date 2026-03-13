package service

import (
	"context"
	"gen/kitex_gen/member"
)

type UpdateProductService struct {
	ctx context.Context
}

// NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *member.UpdateProductReq) (resp *member.ProductResp, err error) {
	// Finish your business logic.

	return
}
