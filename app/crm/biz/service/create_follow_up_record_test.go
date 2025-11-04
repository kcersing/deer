package service

import (
	"context"
	"testing"
)

func TestCreateFollowUpRecord_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateFollowUpRecordService(ctx)
	// init req and assert value

	req := &crm.CreateFollowUpRecordReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
