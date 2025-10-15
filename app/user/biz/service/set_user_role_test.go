package service

import (
	"context"
	User "gen/kitex_gen/user"
	"testing"
)

func TestSetUserRole_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSetUserRoleService(ctx)
	// init req and assert value

	req := &User.SetUserRoleReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
