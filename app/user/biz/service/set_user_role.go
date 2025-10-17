package service

import (
	"context"
	Base "gen/kitex_gen/base"
	User "gen/kitex_gen/user"
	"user/biz/dal/db"
	"user/biz/dal/db/ent/user"
)

type SetUserRoleService struct {
	ctx context.Context
} // NewSetUserRoleService new SetUserRoleService
func NewSetUserRoleService(ctx context.Context) *SetUserRoleService {
	return &SetUserRoleService{ctx: ctx}
}

// Run create note info
func (s *SetUserRoleService) Run(req *User.SetUserRoleReq) (resp *Base.NilResponse, err error) {
	_, err = db.Client.User.Update().
		Where(user.IDEQ(req.GetId())).AddUserRoleIDs(int(req.GetRoleId())).
		Save(s.ctx)

	if err != nil {
		return nil, err
	}
	return
}
