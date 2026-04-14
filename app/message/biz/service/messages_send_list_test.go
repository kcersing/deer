package service

import (
	"context"
	"gen/kitex_gen/message"
	"testing"
)

func TestMessagesSendList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewMessagesSendListService(ctx)
	// init req and assert value

	req := &message.MessagesSendListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
