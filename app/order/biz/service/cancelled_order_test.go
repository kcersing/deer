package service

import (
	"context"
	order "gen/kitex_gen/order"
	"testing"
)

func TestCancelledOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCancelledOrderService(ctx)
	// init req and assert value

	req := &order.CreateOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
