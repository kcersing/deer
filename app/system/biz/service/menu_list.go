package service

import (
	"context"
	system "gen/kitex_gen/system"
)

type MenuListService struct {
	ctx context.Context
} // NewMenuListService new MenuListService
func NewMenuListService(ctx context.Context) *MenuListService {
	return &MenuListService{ctx: ctx}
}

// Run create note info
func (s *MenuListService) Run(req *system.MenuListReq) (resp *system.MenuListResp, err error) {
	// Finish your business logic.

	return
}
