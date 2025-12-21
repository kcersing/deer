package service

import (
	"context"
	base "gen/kitex_gen/base"
	product "gen/kitex_gen/product"
	"product/biz/dal/db"
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

	_, err = db.Client.Product.UpdateOneID(req.GetProductId()).
		SetStock(-req.GetCount()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}

	return
}
