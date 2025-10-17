package service

import (
	"context"
	system "gen/kitex_gen/system"
	"testing"
)

func TestDeleteLog_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteLogService(ctx)
	// init req and assert value

	req := &system.DeleteLog{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
