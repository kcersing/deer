package service

import (
	"context"
	product "gen/kitex_gen/product"
)

type SearchProductService struct {
	ctx context.Context
} // NewSearchProductService new SearchProductService
func NewSearchProductService(ctx context.Context) *SearchProductService {
	return &SearchProductService{ctx: ctx}
}

// Run create note info
func (s *SearchProductService) Run(req *product.SearchProductReq) (resp *product.ProductListResp, err error) {
	// Finish your business logic.

	return
}
