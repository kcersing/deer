package service

import (
	"context"
	"testing"
)

func TestUpdateOpportunities_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateOpportunitiesService(ctx)
	// init req and assert value

	req := &crm.UpdateOpportunitiesReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
