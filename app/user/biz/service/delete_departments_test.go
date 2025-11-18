package service

import (
	"context"
	"gen/kitex_gen/base"
	"testing"
)

func TestDeleteDepartments_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteDepartmentsService(ctx)
	// init req and assert value

	req := &base.IdReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
