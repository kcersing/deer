package service

import (
	"context"
	Member "gen/kitex_gen/member"
	"testing"
)

func TestCreateMember_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateMemberService(ctx)
	// init req and assert value

	req := &Member.CreateMemberReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
