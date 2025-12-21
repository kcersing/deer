package service

import (
	"context"
	base "gen/kitex_gen/base"
	product "gen/kitex_gen/product"
	"product/biz/dal/db"
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

	_, err = db.Client.Product.UpdateOneID(req.GetProductId()).
		SetStock(+req.GetCount()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
