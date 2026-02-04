package service

import (
	"context"
	base "message/1/kitex_gen/base"
	message "message/1/kitex_gen/message"
	"testing"
)

func TestUpdateSend_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateSendService(ctx)
	// init req and assert value

	req := &message.SendMessagesReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
