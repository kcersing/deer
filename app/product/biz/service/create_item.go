package service

import (
	"context"
	product "gen/kitex_gen/product"
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

	save, err := db.Client.ProductItem.Create().
		SetName(req.GetName()).
		SetPic(req.GetPic()).
		SetDesc(req.GetDesc()).
		SetType(req.GetType()).
		SetDuration(req.GetDuration()).
		SetLength(req.GetLength()).
		SetCount(req.GetCount()).
		SetPrice(req.GetPrice()).
		SetActiveAt(req.GetActiveAt()).
		SetExpiredAt(req.GetExpiredAt()).
		SetTagID(req.GetTagID()).
		Save()
	return
}
