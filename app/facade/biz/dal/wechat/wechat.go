package wechat

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/medivhzhan/weapp/v2"
)

type AuthServiceImpl struct {
	AppId     string
	AppSecret string
}

func (s *AuthServiceImpl) Resolve(code string) string {
	resp, err := weapp.Login(s.AppId, s.AppSecret, code)
	if err != nil {
		klog.Errorf("WeApp.Login Err: %v code:%s", err, code)
		return ""
	}
	if err := resp.GetResponseError(); err != nil {
		klog.Errorf("WeApp.Login resp Err: %v code:%s", err, code)
		return ""
	}
	return resp.OpenID
}
