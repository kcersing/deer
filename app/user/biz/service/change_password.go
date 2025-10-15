package service

import (
	"common/pkg/encrypt"
	"context"
	Base "gen/kitex_gen/base"
	User "gen/kitex_gen/user"
	"user/biz/dal/db"
	"user/biz/dal/db/ent/user"
)

type ChangePasswordService struct {
	ctx context.Context
} // NewChangePasswordService new ChangePasswordService
func NewChangePasswordService(ctx context.Context) *ChangePasswordService {
	return &ChangePasswordService{ctx: ctx}
}

// Run create note info
func (s *ChangePasswordService) Run(req *User.ChangePasswordReq) (resp *Base.NilResponse, err error) {
	password, _ := encrypt.Crypt(req.GetPassword())
	_, err = db.Client.User.Update().Where(user.IDEQ(req.GetId())).SetPassword(password).Save(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
