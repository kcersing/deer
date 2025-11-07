package service

import (
	"common/pkg/errno"
	"context"
	User "gen/kitex_gen/user"
	"user/biz/dal/db"
	"user/biz/dal/db/ent/user"
)

type CreateUserService struct {
	ctx context.Context
} // NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// Run create note info
func (s *CreateUserService) Run(req *User.CreateUserReq) (resp *User.UserResp, err error) {

	ok, _ := db.Client.User.Query().Where(user.UsernameEQ(req.GetUsername())).Exist(s.ctx)
	if ok {
		return nil, errno.UserAlreadyExistErr
	}
	//ok, _ = db.Client.User.Query().Where(user.MobileEQ(*req.Username)).Exist(s.ctx)
	//if ok {
	//	return nil, errno.UserMobileExistErr
	//}
	_, err = db.Client.User.Create().
		SetUsername(req.GetUsername()).
		SetPassword(req.GetPassword()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	resp = &User.UserResp{
		Data: &User.User{
			Username: req.GetUsername(),
		},
	}
	return
}
