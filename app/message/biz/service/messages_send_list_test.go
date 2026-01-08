package service

import (
	"context"
	message "gen/kitex_gen/message"
	"testing"
)

func TestMessagesSendList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewMessagesSendListService(ctx)
	// init req and assert value

	req := &message.MessagesListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
