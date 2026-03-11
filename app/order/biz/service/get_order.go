package service

import (
	"common/pkg/errno"
	"context"
	"gen/kitex_gen/base"
	order "gen/kitex_gen/order"
	"order/biz/infras/repo"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
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

	orderFromDB, err := repo.NewOrderRepo().FindById(s.ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	//klog.Info(orderFromDB)

	var dto base.Order
	if err = copier.CopyWithOption(&dto, orderFromDB, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		klog.Errorf("failed to copy entity to DTO: %w", err)
		return nil, err
	}

	resp = &order.OrderResp{
		Data: &dto,
	}

	return
}
