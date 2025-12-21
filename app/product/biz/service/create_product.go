package service

import (
	"common/pkg/errno"
	"common/pkg/utils"
	"context"
	product "gen/kitex_gen/product"
	"product/biz/convert"
	"product/biz/dal/db"
)

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.ProductResp, err error) {
	// Finish your business logic.
	signSalesAt, err := utils.GetStringDateTime(req.GetSignSalesAt())
	if err != nil {
		return nil, errno.TimeFormatErr
	}

	endSalesAt, err := utils.GetStringDateTime(req.GetEndSalesAt())
	if err != nil {
		return nil, errno.TimeFormatErr
	}

	entity, err := db.Client.Product.Create().
		SetName(req.GetName()).
		SetCode(req.GetCode()).
		SetPic(req.GetPic()).
		SetDesc(req.GetDesc()).
		SetPrice(req.GetPrice()).
		SetStock(req.GetStock()).
		SetIsSales(req.GetIsSales()).
		SetSignSalesAt(signSalesAt).
		SetEndSalesAt(endSalesAt).
		SetCreatedID(req.GetCreatedId()).
		SetStatus(req.GetStatus()).
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
