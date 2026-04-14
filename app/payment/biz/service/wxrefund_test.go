package service

import (
	"context"

	payment "gen/kitex_gen/payment"
	"testing"
)

func TestWXRefund_Run(t *testing.T) {
	ctx := context.Background()
	s := NewWXRefundService(ctx)
	// init req and assert value

	req := &payment.WXRefundReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
