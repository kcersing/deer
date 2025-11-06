package service

import (
	"context"
	product "gen/kitex_gen/product"
)

type UpdateItemService struct {
	ctx context.Context
} // NewUpdateItemService new UpdateItemService
func NewUpdateItemService(ctx context.Context) *UpdateItemService {
	return &UpdateItemService{ctx: ctx}
}

// Run create note info
func (s *UpdateItemService) Run(req *product.UpdateItemReq) (resp *product.ItemResp, err error) {
	// Finish your business logic.

	return
}
