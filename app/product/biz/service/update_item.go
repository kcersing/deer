package service

import (
	"context"
	product "gen/kitex_gen/product"
	"product/biz/convert"
	"product/biz/dal/db"
)

type UpdateItemService struct {
	ctx context.Context
} // NewUpdateItemService new UpdateItemService
func NewUpdateItemService(ctx context.Context) *UpdateItemService {
	return &UpdateItemService{ctx: ctx}
}

// Run create note info
func (s *UpdateItemService) Run(req *product.UpdateItemReq) (resp *product.ItemResp, err error) {
	// Finish your business logic.

	entity, err := db.Client.Item.UpdateOneID(req.GetId()).
		SetName(req.GetName()).
		SetPic(req.GetPic()).
		SetDesc(req.GetDesc()).
		SetType(req.GetType()).
		SetDuration(req.GetDuration()).
		SetLength(req.GetLength()).
		SetCount(req.GetCount()).
		SetPrice(req.GetPrice()).
		SetTagID(req.GetTagId()).
		SetStatus(req.GetStatus()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	dataResp := convert.EntToItem(entity)
	resp = &product.ItemResp{
		Data: dataResp,
	}
	return
}
