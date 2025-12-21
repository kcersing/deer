package service

import (
	"context"
	base "gen/kitex_gen/base"
	"product/biz/dal/db"
	"product/biz/dal/db/ent/product"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	_, err = db.Client.Product.Delete().Where(product.IDEQ(req.GetId())).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return

}
