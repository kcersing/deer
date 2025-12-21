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

type GetItemService struct {
	ctx context.Context
} // NewGetItemService new GetItemService
func NewGetItemService(ctx context.Context) *GetItemService {
	return &GetItemService{ctx: ctx}
}

// Run create note info
func (s *GetItemService) Run(req *base.IdReq) (resp *product.ItemResp, err error) {
	// Finish your business logic.

	if req == nil {
		return nil, errno.InvalidParameterErr
	}

	entity, err := db.Client.Item.Get(s.ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errno.NotFound
		}
		return nil, errno.QueryFailed
	}
	dataResp := convert.EntToItem(entity)
	resp = &product.ItemResp{
		Data: dataResp,
	}
	return
}
