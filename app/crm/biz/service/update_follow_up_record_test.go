package service

import (
	"context"
	"testing"
)

func TestUpdateFollowUpRecord_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateFollowUpRecordService(ctx)
	// init req and assert value

	req := &crm.UpdateFollowUpRecordReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
