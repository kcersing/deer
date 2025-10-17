package service

import (
	"context"
	system "gen/kitex_gen/system"
)

type CreateRoleService struct {
	ctx context.Context
} // NewCreateRoleService new CreateRoleService
func NewCreateRoleService(ctx context.Context) *CreateRoleService {
	return &CreateRoleService{ctx: ctx}
}

// Run create note info
func (s *CreateRoleService) Run(req *system.CreateRoleReq) (resp *system.RoleResp, err error) {
	// Finish your business logic.

	return
}
