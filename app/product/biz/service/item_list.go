package service

import (
	"context"
	product "gen/kitex_gen/product"
)

type ItemListService struct {
	ctx context.Context
} // NewItemListService new ItemListService
func NewItemListService(ctx context.Context) *ItemListService {
	return &ItemListService{ctx: ctx}
}

// Run create note info
func (s *ItemListService) Run(req *product.ItemListReq) (resp *product.ItemListResp, err error) {
	// Finish your business logic.

	return
}
