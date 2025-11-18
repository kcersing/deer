package service

import (
	"context"
	"gen/kitex_gen/user"
	"testing"
)

func TestGetDepartmentsList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetDepartmentsListService(ctx)
	// init req and assert value

	req := &user.GetDepartmentsListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
