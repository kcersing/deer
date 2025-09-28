package service

import (
	"context"
	base "order/kitex_gen/deer/base"
	"testing"
)

func TestDeleteOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteOrderService(ctx)
	// init req and assert value

	req := &base.IDReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
