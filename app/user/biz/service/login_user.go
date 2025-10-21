package service

import (
	"common/pkg/encrypt"
	"common/pkg/errno"
	"context"
	Base "gen/kitex_gen/base"
	User "gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"user/biz/convert"
	"user/biz/dal/db"
	"user/biz/dal/db/ent/user"
)

type LoginUserService struct {
	ctx context.Context
} // NewLoginUserService new LoginUserService
func NewLoginUserService(ctx context.Context) *LoginUserService {
	return &LoginUserService{ctx: ctx}
}

// Run create note info
func (s *LoginUserService) Run(req *Base.CheckAccountReq) (resp *User.UserResp, err error) {
	// Finish your business logic.

	only, err := db.Client.User.Query().Where(user.Username(req.GetUsername())).Only(s.ctx)
	klog.Info(only)
	if err != nil {
		return nil, errno.UserNotExistErr
	}
	if ok := encrypt.VerifyPassword(req.Password, only.Password); !ok {
		return nil, errno.LoginPasswordErr
	}
	resp.Data = convert.EntToUser(only)
	return
}
