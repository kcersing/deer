package service

import (
	"context"

	payment "gen/kitex_gen/payment"
	"testing"
)

func TestWXQRPay_Run(t *testing.T) {
	ctx := context.Background()
	s := NewWXQRPayService(ctx)
	// init req and assert value

	req := &payment.WXQRPayReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
