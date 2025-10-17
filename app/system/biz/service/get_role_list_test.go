package service

import (
	"context"
	system "gen/kitex_gen/system"
	"testing"
)

func TestGetRoleList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetRoleListService(ctx)
	// init req and assert value

	req := &system.GetRoleListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
