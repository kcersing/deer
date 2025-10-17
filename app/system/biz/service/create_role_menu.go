package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
)

type CreateRoleMenuService struct {
	ctx context.Context
} // NewCreateRoleMenuService new CreateRoleMenuService
func NewCreateRoleMenuService(ctx context.Context) *CreateRoleMenuService {
	return &CreateRoleMenuService{ctx: ctx}
}

// Run create note info
func (s *CreateRoleMenuService) Run(req *system.CreateMenuAuthReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
