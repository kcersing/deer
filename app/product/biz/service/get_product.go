package service

import (
	"common/pkg/errno"
	"context"
	base "gen/kitex_gen/base"
	product "gen/kitex_gen/product"
	"product/biz/convert"
	"product/biz/dal/db"
	"product/biz/dal/db/ent"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *base.IdReq) (resp *product.ProductResp, err error) {
	// Finish your business logic.
	if req == nil {
		return nil, errno.InvalidParameterErr
	}

	entity, err := db.Client.Product.Get(s.ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errno.NotFound
		}
		return nil, errno.QueryFailed
	}
	dataResp := convert.EntToProduct(entity)
	resp = &product.ProductResp{
		Data: dataResp,
	}
	return

}
