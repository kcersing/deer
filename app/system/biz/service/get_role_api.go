package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
)

type GetRoleApiService struct {
	ctx context.Context
} // NewGetRoleApiService new GetRoleApiService
func NewGetRoleApiService(ctx context.Context) *GetRoleApiService {
	return &GetRoleApiService{ctx: ctx}
}

// Run create note info
func (s *GetRoleApiService) Run(req *base.IdReq) (resp *system.MenuListResp, err error) {
	// Finish your business logic.

	return
}
