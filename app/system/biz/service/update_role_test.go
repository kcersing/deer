package service

import (
	"context"
	system "gen/kitex_gen/system"
	"testing"
)

func TestUpdateRole_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateRoleService(ctx)
	// init req and assert value

	req := &system.UpdateRoleReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
