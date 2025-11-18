package service

import (
	"context"
	"gen/kitex_gen/user"
	"testing"
)

func TestCreateDepartments_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateDepartmentsService(ctx)
	// init req and assert value

	req := &user.CreateDepartmentsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
