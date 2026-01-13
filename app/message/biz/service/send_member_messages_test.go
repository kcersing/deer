package service

import (
	"context"

	message "gen/kitex_gen/message"
	"testing"
)

func TestSendMemberMessages_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendMemberMessagesService(ctx)
	// init req and assert value

	req := &message.SendMemberMessagesReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
