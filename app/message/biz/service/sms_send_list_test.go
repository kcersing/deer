package service

import (
	"context"
	message "message/1/kitex_gen/message"
	"testing"
)

func TestSmsSendList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSmsSendListService(ctx)
	// init req and assert value

	req := &message.SmsSendListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
