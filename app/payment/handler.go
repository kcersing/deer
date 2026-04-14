package main

import (
	"context"
	base "gen/kitex_gen/base"
	payment "gen/kitex_gen/payment"
	"payment/biz/service"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// WXPay implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) WXPay(ctx context.Context, req *payment.WXPayReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewWXPayService(ctx).Run(req)

	return resp, err
}

// WXQRPay implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) WXQRPay(ctx context.Context, req *payment.WXQRPayReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewWXQRPayService(ctx).Run(req)

	return resp, err
}

// WXRefund implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) WXRefund(ctx context.Context, req *payment.WXRefundReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewWXRefundService(ctx).Run(req)

	return resp, err
}
