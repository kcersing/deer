package service

import (
	"context"
	crm "gen/kitex_gen/crm"
	"testing"
)

func TestOpportunitiesList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewOpportunitiesListService(ctx)
	// init req and assert value

	req := &crm.OpportunitiesListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
