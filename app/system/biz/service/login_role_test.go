package service

import (
	"context"
	base "gen/kitex_gen/base"
	"testing"
)

func TestLoginRole_Run(t *testing.T) {
	ctx := context.Background()
	s := NewLoginRoleService(ctx)
	// init req and assert value

	req := &base.CheckAccountReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
