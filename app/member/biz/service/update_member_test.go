package service

import (
	"context"
	Member "gen/kitex_gen/member"
	"testing"
)

func TestUpdateOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateMemberService(ctx)
	// init req and assert value

	req := &Member.UpdateMemberReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
