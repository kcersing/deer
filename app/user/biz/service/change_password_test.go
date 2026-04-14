package service

import (
	"common/pkg/encrypt"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
)

func TestChangePassword_Run(t *testing.T) {
	//ctx := context.Background()
	//s := NewChangePasswordService(ctx)
	//// init req and assert value
	//
	//req := &User.ChangePasswordReq{}
	//resp, err := s.Run(req)
	//t.Logf("err: %v", err)
	//t.Logf("resp: %v", resp)

	password, _ := encrypt.Crypt("123456")

	klog.Info(password)

	// todo: edit your unit test

}
