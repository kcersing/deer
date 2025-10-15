package service

import (
	"context"
	Base "gen/kitex_gen/base"
	User "gen/kitex_gen/user"
)

type SetUserRoleService struct {
	ctx context.Context
} // NewSetUserRoleService new SetUserRoleService
func NewSetUserRoleService(ctx context.Context) *SetUserRoleService {
	return &SetUserRoleService{ctx: ctx}
}

// Run create note info
func (s *SetUserRoleService) Run(req *User.SetUserRoleReq) (resp *Base.NilResponse, err error) {
	return
}
