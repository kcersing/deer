package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
)

type LoginRoleService struct {
	ctx context.Context
} // NewLoginRoleService new LoginRoleService
func NewLoginRoleService(ctx context.Context) *LoginRoleService {
	return &LoginRoleService{ctx: ctx}
}

// Run create note info
func (s *LoginRoleService) Run(req *base.CheckAccountReq) (resp *system.RoleResp, err error) {
	// Finish your business logic.

	return
}
