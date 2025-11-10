package service

import (
	"context"
	Base "gen/kitex_gen/base"
	"testing"
)

func TestMenu_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetMenuService(ctx)
	// init req and assert value

	req := &Base.IdReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
