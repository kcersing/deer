package service

import (
	"context"
	system "gen/kitex_gen/system"
	"testing"
)

func TestLogList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewLogListService(ctx)
	// init req and assert value

	req := &system.LogListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
