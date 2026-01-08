package service

import (
	"context"
	base "gen/kitex_gen/base"
	message "gen/kitex_gen/message"
	"testing"
)

func TestSendUserMessages_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendUserMessagesService(ctx)
	// init req and assert value

	req := &message.SendUserMessagesReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
