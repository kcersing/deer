package service

import (
	"common/pkg/errno"
	"context"
	order "gen/kitex_gen/order"
	"order/biz/infras/repo"

	"github.com/cloudwego/kitex/pkg/klog"
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
	klog.Info(req)
	if req == nil {
		return nil, errno.InvalidParameterErr
	}

	orderFromDB, err := repo.NewOrderRepo(nil).FindById(s.ctx, req.GetId())
	klog.Info(orderFromDB)

	//if err != nil {
	//	if ent.IsNotFound(err) {
	//		return nil, errno.NotFound
	//	}
	//	return nil, errno.QueryFailed
	//}

	//klog.Info(orderFromDB)
	//
	resp = &order.OrderResp{
		Data: nil,
	}

	return
}
