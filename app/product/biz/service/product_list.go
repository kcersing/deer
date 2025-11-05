package service

import (
	"context"
	product "gen/kitex_gen/product"
)

type ProductListService struct {
	ctx context.Context
} // NewProductListService new ProductListService
func NewProductListService(ctx context.Context) *ProductListService {
	return &ProductListService{ctx: ctx}
}

// Run create note info
func (s *ProductListService) Run(req *product.ListReq) (resp *product.ProductListResp, err error) {
	// Finish your business logic.

	return
}
