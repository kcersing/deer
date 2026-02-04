package service

import (
	"context"
	message "message/1/kitex_gen/message"
	"testing"
)

func TestSendSms_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendSmsService(ctx)
	// init req and assert value

	req := &message.SendSmsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
