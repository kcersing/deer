package service

import (
	"common/pkg/errno"
	"context"
	User "gen/kitex_gen/user"
	"time"
	"user/biz/convert"
	"user/biz/dal/db"
	"user/biz/dal/db/ent/user"
)

type UpdateUserService struct {
	ctx context.Context
} // NewUpdateUserService new UpdateUserService
func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

// Run Update note info
func (s *UpdateUserService) Run(req *User.UpdateUserReq) (resp *User.UserResp, err error) {

	ok, _ := db.Client.User.Query().Where(user.MobileEQ(req.GetMobile()), user.IDNEQ(req.GetId())).Exist(s.ctx)
	if ok {
		return nil, errno.UserMobileExistErr
	}

	birthday, err := time.Parse(time.DateOnly, req.GetBirthday())
	if err != nil {
		return nil, errno.TimeFormatErr
	}
	save, err := db.Client.User.UpdateOneID(req.GetId()).
		SetAvatar(req.GetAvatar()).
		SetMobile(req.GetMobile()).
		SetName(req.GetName()).
		SetStatus(req.GetStatus()).
		SetGender(req.GetGender()).
		SetBirthday(birthday).
		SetDesc(req.GetDesc()).
		//SetRoleID(req.GetRoleId()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	resp = &User.UserResp{
		Data: convert.EntToUser(save),
	}
	return
}
