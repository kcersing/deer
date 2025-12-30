package service

import (
	"context"
	"gen/kitex_gen/base"
	product "gen/kitex_gen/product"
	"product/biz/dal/db/ent/productitem"
	"user/biz/dal/db/ent/position"

	"product/biz/convert"
	"product/biz/dal/db"
	"product/biz/dal/db/ent"

	"product/biz/dal/db/ent/predicate"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type ItemListService struct {
	ctx context.Context
} // NewItemListService new ItemListService
func NewItemListService(ctx context.Context) *ItemListService {
	return &ItemListService{ctx: ctx}
}

// Run create note info
func (s *ItemListService) Run(req *product.ItemListReq) (resp *product.ItemListResp, err error) {
	// Finish your business logic.

	var (
		dataResp []*base.Item
	)

	var predicates []predicate.ProductItem
	if req.GetName() != "" {
		predicates = append(predicates, productitem.NameContains(req.GetName()))
	}
	if len(req.GetStatus()) > 0 {
		predicates = append(predicates, productitem.StatusIn(req.GetStatus()...))
	}
	if req.GetType() != "" {
		predicates = append(predicates, productitem.TypeEQ(req.GetType()))
	}
	predicates = append(predicates, productitem.DeleteEQ(0))

	all, err := db.Client.Debug().ProductItem.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(position.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)
	hlog.Info(all)
	if err != nil {
		return nil, err
	}
	for _, v := range all {
		dataResp = append(dataResp, convert.EntToItem(v))
	}

	return &product.ItemListResp{
		Data: dataResp,
	}, nil

}
