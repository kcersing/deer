package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
)

type CreateRoleApiService struct {
	ctx context.Context
} // NewCreateRoleApiService new CreateRoleApiService
func NewCreateRoleApiService(ctx context.Context) *CreateRoleApiService {
	return &CreateRoleApiService{ctx: ctx}
}

// Run create note info
func (s *CreateRoleApiService) Run(req *system.CreateMenuAuthReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
