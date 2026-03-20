package service

import (
	"common/pkg/utils"
	"context"

	"gen/kitex_gen/base"
	"gen/kitex_gen/member"
	"gen/kitex_gen/product"
	"member/biz/dal/db"
	"member/rpc/client"

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

	if !(len(req.Items) > 0) {
		return nil, errors.New("require at least one item")
	}
	for _, i := range req.Items {
		data, err := client.ProductClient.GetProduct(s.ctx, &base.IdReq{Id: i.GetProductId()})
		if err != nil {
			return nil, err
		}
		save, err := tx.MemberProduct.Create().
			SetSn(utils.CreateSn()).
			SetCreatedID(req.GetUserId()).
			SetName(i.GetName()).
			SetMemberID(req.GetMemberId()).
			SetOrderID(req.GetOrderId()).
			SetPrice(i.GetPrice()).
			SetProductID(i.GetProductId()).
			SetActual(req.GetActual()).
			Save(s.ctx)
		if err != nil {
			return nil, rollback(tx, errors.Wrap(err, "create Member Product failed"))
		}

		if len(data.Data.Items) > 0 {
			items, err := client.ProductClient.ItemList(s.ctx, &product.ItemListReq{
				Page:     1,
				PageSize: int64(len(data.Data.Items)),
				Ids:      data.Data.Items,
			})
			if err != nil {
				return nil, err
			}

			for _, v := range items.Data {
				_, err = tx.MemberProductProperty.Create().
					SetSn(utils.CreateSn()).
					SetCreatedID(req.GetUserId()).
					SetName(v.Name).
					SetMemberID(req.GetMemberId()).
					SetPrice(v.Price).
					SetType(v.Type).
					SetDuration(v.Duration).
					SetMemberProductID(save.ID).
					SetLength(v.Length).
					SetCount(v.Count).
					SetPropertyID(v.Id).
					Save(s.ctx)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return nil, err
}
