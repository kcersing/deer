package service

import (
	"context"
	system "gen/kitex_gen/system"
)

type GetRoleListService struct {
	ctx context.Context
} // NewGetRoleListService new GetRoleListService
func NewGetRoleListService(ctx context.Context) *GetRoleListService {
	return &GetRoleListService{ctx: ctx}
}

// Run create note info
func (s *GetRoleListService) Run(req *system.GetRoleListReq) (resp *system.RoleListResp, err error) {
	// Finish your business logic.

	return
}
