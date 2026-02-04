package service

import (
	"context"
	message "message/1/kitex_gen/message"
	"testing"
)

func TestSendMessages_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendMessagesService(ctx)
	// init req and assert value

	req := &message.SendMessagesReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
