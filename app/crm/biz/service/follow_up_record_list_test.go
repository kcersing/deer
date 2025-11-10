package service

import (
	"context"
	"gen/kitex_gen/crm"
	"testing"
)

func TestFollowUpRecordList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewFollowUpRecordListService(ctx)
	// init req and assert value

	req := &crm.FollowUpRecordListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
