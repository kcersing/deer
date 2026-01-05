package service

import (
	"context"
	message "gen/kitex_gen/message"
	"testing"
)

func TestSmsList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSmsListService(ctx)
	// init req and assert value

	req := &message.SendSmsListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
