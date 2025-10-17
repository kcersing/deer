package service

import (
	"context"
	system "gen/kitex_gen/system"
	"testing"
)

func TestApiTree_Run(t *testing.T) {
	ctx := context.Background()
	s := NewApiTreeService(ctx)
	// init req and assert value

	req := &system.ApiListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
