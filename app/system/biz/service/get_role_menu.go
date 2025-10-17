package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
)

type GetRoleMenuService struct {
	ctx context.Context
} // NewGetRoleMenuService new GetRoleMenuService
func NewGetRoleMenuService(ctx context.Context) *GetRoleMenuService {
	return &GetRoleMenuService{ctx: ctx}
}

// Run create note info
func (s *GetRoleMenuService) Run(req *base.IdReq) (resp *system.MenuListResp, err error) {
	// Finish your business logic.

	return
}
