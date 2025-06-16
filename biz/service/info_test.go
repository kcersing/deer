package service

import (
	"context"
	order "kcers-order/kitex_gen/order"
	"testing"
)

func TestInfo_Run(t *testing.T) {
	ctx := context.Background()
	s := NewInfoService(ctx)
	// init req and assert value

	req := &order.Req{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
