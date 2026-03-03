package service

import (
	"common/pkg/errno"
	"context"
	order "gen/kitex_gen/order"
	"order/biz/infras/repo"

	"order/biz/dal/convert"
	"order/biz/dal/db"
	"order/biz/dal/db/ent"
)

type GetOrderService struct {
	ctx context.Context
}

// NewGetOrderService new GetOrderService
func NewGetOrderService(ctx context.Context) *GetOrderService {
	return &GetOrderService{ctx: ctx}
}

// Run create note info
func (s *GetOrderService) Run(req *order.GetOrderReq) (resp *order.OrderResp, err error) {
	// Finish your business logic.
	orderFromDB, err := repo.NewOrderRepo(nil).FindById(req.GetId())
	if err != nil {
		return nil, err
	}

	if req == nil {
		return nil, errno.InvalidParameterErr
	}

	entity, err := db.Client.Order.Get(s.ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errno.NotFound
		}
		return nil, errno.QueryFailed
	}
	dataResp := convert.EntToOrder(entity)
	resp = &order.OrderResp{
		Data: dataResp,
	}

	return
}
