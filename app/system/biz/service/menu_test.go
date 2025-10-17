package service

import (
	"context"
	system "gen/kitex_gen/system"
	"testing"
)

func TestMenu_Run(t *testing.T) {
	ctx := context.Background()
	s := NewMenuService(ctx)
	// init req and assert value

	req := &base.IdReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
