package service

import (
	"context"
	base "gen/kitex_gen/base"
	order "gen/kitex_gen/order"
	"order/biz/infras/repo"
)

type GetOrderInfoService struct {
	ctx context.Context
}

// NewGetOrderInfoService new GetOrderInfoService
func NewGetOrderInfoService(ctx context.Context) *GetOrderInfoService {
	return &GetOrderInfoService{ctx: ctx}
}

// Run create note info
func (s *GetOrderInfoService) Run(req *order.GetOrderInfoReq) (resp *order.OrderResp, err error) {
	// Finish your business logic.
	orderFromDB, err := repo.OrderRepoClient.FindById(req.GetId())
	if err != nil {
		return nil, err
	}
	resp = &order.OrderResp{
		Data: &base.Order{
			MemberId:    orderFromDB.MemberId,
			Items:       orderFromDB.Items,
			Sn:          orderFromDB.Sn,
			TotalAmount: orderFromDB.TotalAmount,
			Status:      orderFromDB.Status,
			Nature:      orderFromDB.Nature,
			//CreatedAt:       orderFromDB.CreatedAt.Format(time.RFC3339),
			//CompletionAt:    orderFromDB.CompletionAt.Format(time.RFC3339),
			//CloseAt:         orderFromDB.CloseAt.Format(time.RFC3339),
			//UpdatedAt:       orderFromDB.UpdatedAt.Format(time.RFC3339),
			CancelledReason: orderFromDB.CancelledReason,
			OrderPays:       orderFromDB.OrderPays,
			OrderRefund:     orderFromDB.OrderRefund,
			Id:              orderFromDB.Id,
			CreatedId:       orderFromDB.CreatedId,
			CreatedName:     orderFromDB.CreatedName,
		},
	}
	return
}
