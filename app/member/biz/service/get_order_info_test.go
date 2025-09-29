package service

import (
	"context"
	order "member/kitex_gen/order"
	"testing"
)

func TestGetOrderInfo_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetOrderInfoService(ctx)
	// init req and assert value

	req := &order.GetOrderInfoReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
