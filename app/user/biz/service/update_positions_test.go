package service

import (
	"context"
	"gen/kitex_gen/user"
	"testing"
)

func TestUpdatePositions_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdatePositionsService(ctx)
	// init req and assert value

	req := &user.UpdatePositionsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
