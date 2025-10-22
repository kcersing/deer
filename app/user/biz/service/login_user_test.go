package service

import (
	"context"
	Base "gen/kitex_gen/base"
	"testing"
)

func TestLoginOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewLoginUserService(ctx)
	// init req and assert value

	req := &Base.CheckAccountReq{
		Username: "admin",
		Password: "123456",
		Captcha:  "1234",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
