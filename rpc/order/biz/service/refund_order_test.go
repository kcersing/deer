package service

import (
	"context"
	order "deer/kitex_gen/deer/order"
	"testing"
)

func TestRefundOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRefundOrderService(ctx)
	// init req and assert value

	req := &order.RefundOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
