package service

import (
	"context"
	base "gen/kitex_gen/base"
	product "gen/kitex_gen/product"
)

type DecrStockService struct {
	ctx context.Context
} // NewDecrStockService new DecrStockService
func NewDecrStockService(ctx context.Context) *DecrStockService {
	return &DecrStockService{ctx: ctx}
}

// Run create note info
func (s *DecrStockService) Run(req *product.DecrStockReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
