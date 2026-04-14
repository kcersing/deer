package service

import (
	"common/pkg/encrypt"
	"common/pkg/errno"
	"context"
	Base "gen/kitex_gen/base"
	User "gen/kitex_gen/user"
	"user/biz/dal/db"
	"user/biz/dal/db/ent/user"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type ChangePasswordService struct {
	ctx context.Context
} // NewChangePasswordService new ChangePasswordService
func NewChangePasswordService(ctx context.Context) *ChangePasswordService {
	return &ChangePasswordService{ctx: ctx}
}

// Run create note info
func (s *ChangePasswordService) Run(req *User.ChangePasswordReq) (resp *Base.NilResponse, err error) {

	only, err := db.Client.User.Query().Where(user.ID(req.GetId())).Only(s.ctx)

	if err != nil {
		return nil, errno.UserNotExistErr
	}

	hlog.Info(req.GetOldPassword())
	if ok := encrypt.VerifyPassword(req.GetOldPassword(), only.Password); !ok {
		return nil, errno.LoginPasswordErr
	}

	password, _ := encrypt.Crypt(req.GetPassword())

	_, err = db.Client.User.Update().Where(user.IDEQ(req.GetId())).SetPassword(password).Save(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
