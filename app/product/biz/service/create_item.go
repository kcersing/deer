package service

import (
	"context"
	product "gen/kitex_gen/product"
)

type CreateItemService struct {
	ctx context.Context
} // NewCreateItemService new CreateItemService
func NewCreateItemService(ctx context.Context) *CreateItemService {
	return &CreateItemService{ctx: ctx}
}

// Run create note info
func (s *CreateItemService) Run(req *product.CreateItemReq) (resp *product.ItemResp, err error) {
	// Finish your business logic.

	return
}
