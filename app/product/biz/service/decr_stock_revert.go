package service

import (
	"context"
	base "gen/kitex_gen/base"
	product "gen/kitex_gen/product"
)

type DecrStockRevertService struct {
	ctx context.Context
} // NewDecrStockRevertService new DecrStockRevertService
func NewDecrStockRevertService(ctx context.Context) *DecrStockRevertService {
	return &DecrStockRevertService{ctx: ctx}
}

// Run create note info
func (s *DecrStockRevertService) Run(req *product.DecrStockReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
