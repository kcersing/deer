package service

import (
	"context"
	product "gen/kitex_gen/product"

	"product/biz/convert"
	"product/biz/dal/db"
)

type CreateItemService struct {
	ctx context.Context
} // NewCreateItemService new CreateItemService
func NewCreateItemService(ctx context.Context) *CreateItemService {
	return &CreateItemService{ctx: ctx}
}

// Run create note info
func (s *CreateItemService) Run(req *product.CreateItemReq) (resp *product.ItemResp, err error) {
	// Finish your business logic.

	entity, err := db.Client.Item.Create().
		SetName(req.GetName()).
		SetCode(req.GetCode()).
		SetPic(req.GetPic()).
		SetDesc(req.GetDesc()).
		SetType(req.GetType()).
		SetDuration(req.GetDuration()).
		SetLength(req.GetLength()).
		SetCount(req.GetCount()).
		SetPrice(req.GetPrice()).
		SetTagID(req.GetTagId()).
		SetCreatedID(req.GetCreatedId()).
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
