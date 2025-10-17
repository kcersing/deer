package service

import (
	"context"
	system "gen/kitex_gen/system"
)

type UpdateRoleService struct {
	ctx context.Context
} // NewUpdateRoleService new UpdateRoleService
func NewUpdateRoleService(ctx context.Context) *UpdateRoleService {
	return &UpdateRoleService{ctx: ctx}
}

// Run create note info
func (s *UpdateRoleService) Run(req *system.UpdateRoleReq) (resp *system.RoleResp, err error) {
	// Finish your business logic.

	return
}
