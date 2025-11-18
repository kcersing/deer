package service

import (
	"context"
	"gen/kitex_gen/base"
	"testing"
)

func TestGetDepartments_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetDepartmentsService(ctx)
	// init req and assert value

	req := &base.IdReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
