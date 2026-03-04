package service

import (
	"context"
	"gen/kitex_gen/base"

	order "gen/kitex_gen/order"
	"order/biz/dal/convert"
	"order/biz/dal/db"
	"order/biz/dal/db/ent"
	order2 "order/biz/dal/db/ent/order"
	"order/biz/dal/db/ent/predicate"
)

type GetOrderListService struct {
	ctx context.Context
}

// NewGetOrderListService new GetOrderListService
func NewGetOrderListService(ctx context.Context) *GetOrderListService {
	return &GetOrderListService{ctx: ctx}
}

// Run create note info
func (s *GetOrderListService) Run(req *order.GetOrderListReq) (resp *order.GetOrderListResp, err error) {
	// Finish your business logic.

	var predicates []predicate.Order
	all, err := db.Client.Order.Query().
		WithItems().
		Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(order2.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return nil, err
	}

	dataResp := make([]*base.Order, 0, len(all))

	for _, e := range all {
		dataResp = append(dataResp, convert.EntToOrder(e))
	}

	return &order.GetOrderListResp{
		Data: dataResp,
	}, nil
}
