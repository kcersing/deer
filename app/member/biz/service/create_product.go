package service

import (
	"common/pkg/utils"
	"context"
	"gen/kitex_gen/base"
	"gen/kitex_gen/member"
	product2 "gen/kitex_gen/product"
	"member/biz/dal/db"
	"member/rpc/client"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

type CreateProductService struct {
	ctx context.Context
}

// NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *member.CreateProductReq) (resp *member.ProductResp, err error) {
	// Finish your business logic.

	tx, err := db.Client.Tx(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "starting a transaction:")
	}

	save, err := tx.MemberProduct.Create().
		SetSn(utils.CreateSn()).
		SetCreatedID(req.GetUserId()).
		SetName(req.GetName()).
		SetMemberID(req.GetMemberId()).
		SetOrderID(req.GetOrderId()).
		SetPrice(req.GetPrice()).
		SetProductID(req.GetProductId()).
		Save(s.ctx)
	if err != nil {
		return nil, rollback(tx, errors.Wrap(err, "create Member Product failed"))
	}
	klog.Info(save)
	product, err := client.ProductClient.GetProduct(s.ctx, &base.IdReq{Id: req.GetProductId()})
	if err != nil {
		return nil, err
	}
	productItmes, err := client.ProductClient.ItemList(s.ctx, &product2.ItemListReq{
		Page:     1,
		PageSize: int64(len(product.Data.Items)),
		Ids:      product.Data.Items,
	})
	if err != nil {
		return nil, err
	}
	klog.Info(product)
	klog.Info(productItmes)

	return
}
