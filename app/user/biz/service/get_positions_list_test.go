package service

import (
	"context"
	"gen/kitex_gen/user"
	"testing"
)

func TestGetPositionsList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetPositionsListService(ctx)
	// init req and assert value

	req := &user.GetPositionsListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
