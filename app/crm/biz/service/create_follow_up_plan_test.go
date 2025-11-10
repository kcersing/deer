package service

import (
	"context"
	"gen/kitex_gen/crm"
	"testing"
)

func TestCreateFollowUpPlan_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateFollowUpPlanService(ctx)
	// init req and assert value

	req := &crm.CreateFollowUpPlanReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
