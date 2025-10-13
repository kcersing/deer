package service

import (
	"context"
	order "gen/kitex_gen/user"
	"testing"
)

func TestCreateOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateUserService(ctx)
	// init req and assert value

	req := &order.GetUserListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
