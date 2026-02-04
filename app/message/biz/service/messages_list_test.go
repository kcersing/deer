package service

import (
	"context"
	message "message/1/kitex_gen/message"
	"testing"
)

func TestMessagesList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewMessagesListService(ctx)
	// init req and assert value

	req := &message.MessagesListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
