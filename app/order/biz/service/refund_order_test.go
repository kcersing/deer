package service

import (
	"context"
	order "gen/kitex_gen/order"
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
