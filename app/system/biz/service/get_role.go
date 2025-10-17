package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
)

type GetRoleService struct {
	ctx context.Context
} // NewGetRoleService new GetRoleService
func NewGetRoleService(ctx context.Context) *GetRoleService {
	return &GetRoleService{ctx: ctx}
}

// Run create note info
func (s *GetRoleService) Run(req *base.IdReq) (resp *system.RoleResp, err error) {
	// Finish your business logic.

	return
}
