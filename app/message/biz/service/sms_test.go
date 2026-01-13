package service

import (
	"context"
	base "gen/kitex_gen/base"

	"testing"
)

func TestSms_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSmsService(ctx)
	// init req and assert value

	req := &base.IdReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
