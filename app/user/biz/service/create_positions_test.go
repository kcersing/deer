package service

import (
	"context"
	"gen/kitex_gen/user"
	"testing"
)

func TestCreatePositions_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreatePositionsService(ctx)
	// init req and assert value

	req := &user.CreatePositionsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
