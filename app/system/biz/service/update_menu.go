package service

import (
	"context"
	system "gen/kitex_gen/system"
)

type UpdateMenuService struct {
	ctx context.Context
} // NewUpdateMenuService new UpdateMenuService
func NewUpdateMenuService(ctx context.Context) *UpdateMenuService {
	return &UpdateMenuService{ctx: ctx}
}

// Run create note info
func (s *UpdateMenuService) Run(req *system.UpdateMenuReq) (resp *system.MenuResp, err error) {
	// Finish your business logic.

	return
}
