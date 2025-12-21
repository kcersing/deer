package service

import (
	"context"
	base "gen/kitex_gen/base"
	product "gen/kitex_gen/product"
)

type GetItemService struct {
	ctx context.Context
} // NewGetItemService new GetItemService
func NewGetItemService(ctx context.Context) *GetItemService {
	return &GetItemService{ctx: ctx}
}

// Run create note info
func (s *GetItemService) Run(req *base.IdReq) (resp *product.ItemResp, err error) {
	// Finish your business logic.

	return
}
