package service

import (
	"context"
	"gen/kitex_gen/crm"
	"testing"
)

func TestCreateOpportunities_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateOpportunitiesService(ctx)
	// init req and assert value

	req := &crm.CreateOpportunitiesReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
