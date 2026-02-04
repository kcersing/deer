package service

import (
	"context"
	base "message/1/kitex_gen/base"
	"testing"
)

func TestDeleteMessages_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteMessagesService(ctx)
	// init req and assert value

	req := &base.IdReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
