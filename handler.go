package main

import (
	"context"
	"kcers-order/biz/service"
	order "kcers-order/kitex_gen/order"
)

// OrderImpl implements the last service interface defined in the IDL.
type OrderImpl struct{}

// Info implements the OrderImpl interface.
func (s *OrderImpl) Info(ctx context.Context, req *order.Req) (resp *order.Resp, err error) {
	resp, err = service.NewInfoService(ctx).Run(req)

	return resp, err
}
