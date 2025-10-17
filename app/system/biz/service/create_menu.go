package service

import (
	"context"
	system "gen/kitex_gen/system"
)

type CreateMenuService struct {
	ctx context.Context
} // NewCreateMenuService new CreateMenuService
func NewCreateMenuService(ctx context.Context) *CreateMenuService {
	return &CreateMenuService{ctx: ctx}
}

// Run create note info
func (s *CreateMenuService) Run(req *system.CreateMenuReq) (resp *system.MenuResp, err error) {
	// Finish your business logic.

	return
}
