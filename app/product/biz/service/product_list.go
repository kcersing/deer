package service

import (
	"context"
	"gen/kitex_gen/base"
	product "gen/kitex_gen/product"
	"product/biz/convert"
	"product/biz/dal/db"
	"product/biz/dal/db/ent"
	"product/biz/dal/db/ent/predicate"
	product2 "product/biz/dal/db/ent/product"
	"user/biz/dal/db/ent/position"
)

type ProductListService struct {
	ctx context.Context
} // NewProductListService new ProductListService
func NewProductListService(ctx context.Context) *ProductListService {
	return &ProductListService{ctx: ctx}
}

// Run create note info
func (s *ProductListService) Run(req *product.ListReq) (resp *product.ProductListResp, err error) {
	// Finish your business logic.

	var (
		dataResp []*base.Product
	)

	var predicates []predicate.Product
	if req.GetName() != "" {
		predicates = append(predicates, product2.NameContains(req.GetName()))
	}
	if req.GetStatus() != nil {
		predicates = append(predicates, product2.StatusIn(req.GetStatus()...))
	}

	all, err := db.Client.Product.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(position.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)

	if err != nil {
		return nil, err
	}
	for _, v := range all {
		dataResp = append(dataResp, convert.EntToProduct(v))
	}

	return &product.ProductListResp{
		Data: dataResp,
	}, nil
}
