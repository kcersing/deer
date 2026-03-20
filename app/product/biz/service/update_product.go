package service

import (
	"common/pkg/errno"
	"common/pkg/utils"
	"context"
	product "gen/kitex_gen/product"
	"product/biz/convert"
	"product/biz/dal/db"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.ProductResp, err error) {
	// Finish your business logic.

	signSalesAt, err := utils.GetStringDateOnlyZeroTime(req.GetSignSalesAt())
	if err != nil {
		return nil, errno.TimeFormatErr
	}

	endSalesAt, err := utils.GetStringDateOnlyZeroTime(req.GetEndSalesAt())
	if err != nil {
		return nil, errno.TimeFormatErr
	}

	entity, err := db.Client.Product.UpdateOneID(req.CreatedId).
		SetName(req.GetName()).
		SetCode(req.GetCode()).
		SetPic(req.GetPic()).
		SetDesc(req.GetDesc()).
		SetPrice(req.GetPrice() * 100).
		SetStock(req.GetStock()).
		SetIsSales(req.GetIsSales()).
		SetSignSalesAt(signSalesAt).
		SetEndSalesAt(endSalesAt).
		SetCreatedID(req.GetCreatedId()).
		SetStatus(req.GetStatus()).
		SetItems(req.GetItems()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	dataResp := convert.EntToProduct(entity)
	resp = &product.ProductResp{
		Data: dataResp,
	}
	return
}
