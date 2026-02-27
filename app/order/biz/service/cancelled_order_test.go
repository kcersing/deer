package service

import (
	"context"
	"gen/kitex_gen/order"
	"testing"
)

func TestCancelledOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCancelledOrderService(ctx)
	// init req and assert value

	req := &order.CancelledOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
