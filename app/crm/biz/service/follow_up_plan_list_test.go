package service

import (
	"context"
	"gen/kitex_gen/crm"
	"testing"
)

func TestFollowUpPlanList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewFollowUpPlanListService(ctx)
	// init req and assert value

	req := &crm.FollowUpPlanListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
