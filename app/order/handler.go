package main

import (
	"context"
	base "gen/kitex_gen/base"
	order "gen/kitex_gen/order"
	"order/biz/service"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// GetOrderInfo implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrderInfo(ctx context.Context, req *order.GetOrderInfoReq) (resp *order.OrderResp, err error) {
	resp, err = service.NewGetOrderInfoService(ctx).Run(req)

	return resp, err
}

// GetOrderList implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrderList(ctx context.Context, req *order.GetOrderListReq) (resp *order.GetOrderListResp, err error) {
	resp, err = service.NewGetOrderListService(ctx).Run(req)

	return resp, err
}

// DeleteOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) DeleteOrder(ctx context.Context, req *base.IDReq) (resp *base.BaseResp, err error) {
	resp, err = service.NewDeleteOrderService(ctx).Run(req)

	return resp, err
}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order.GetOrderListReq) (resp *order.OrderResp, err error) {
	resp, err = service.NewCreateOrderService(ctx).Run(req)

	return resp, err
}

// Payment implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) Payment(ctx context.Context, req *order.PaymentReq) (resp *order.OrderResp, err error) {
	resp, err = service.NewPaymentService(ctx).Run(req)

	return resp, err
}

// CancelledOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CancelledOrder(ctx context.Context, req *order.CreateOrderReq) (resp *base.BaseResp, err error) {
	resp, err = service.NewCancelledOrderService(ctx).Run(req)

	return resp, err
}

// RefundOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) RefundOrder(ctx context.Context, req *order.RefundOrderReq) (resp *base.BaseResp, err error) {
	resp, err = service.NewRefundOrderService(ctx).Run(req)

	return resp, err
}
