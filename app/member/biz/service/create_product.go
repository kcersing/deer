package service

import (
	"context"
	"errors"
	"gen/kitex_gen/base"
	"gen/kitex_gen/member"
	"gen/kitex_gen/product"
	"member/rpc/client"

	"github.com/cloudwego/kitex/pkg/klog"
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

	klog.Info("CreateProductService Run: ", req)
	//tx, err := db.Client.Tx(s.ctx)
	//if err != nil {
	//	return nil, errors.Wrap(err, "starting a transaction:")
	//}

	if !(len(req.Items) > 0) {
		return nil, errors.New("require at least one item")
	}
	for _, i := range req.Items {
		data, err := client.ProductClient.GetProduct(s.ctx, &base.IdReq{Id: i.GetProductId()})
		if err != nil {
			return nil, err
		}

		klog.Info("product Run: ", data)
		itmes, err := client.ProductClient.ItemList(s.ctx, &product.ItemListReq{
			Page:     1,
			PageSize: int64(len(data.Data.Items)),
			Ids:      data.Data.Items,
		})
		if err != nil {
			return nil, err
		}
		klog.Info("productItmes Run: ", itmes)
	}

	//save, err := tx.MemberProduct.Create().
	//	SetSn(utils.CreateSn()).
	//	SetCreatedID(req.GetUserId()).
	//	SetName(req.GetName()).
	//	SetMemberID(req.GetMemberId()).
	//	SetOrderID(req.GetOrderId()).
	//	SetPrice(req.GetPrice()).
	//	SetProductID(req.GetProductId()).
	//	Save(s.ctx)
	//if err != nil {
	//	return nil, rollback(tx, errors.Wrap(err, "create Member Product failed"))
	//}
	//klog.Info(save)

	return
}
