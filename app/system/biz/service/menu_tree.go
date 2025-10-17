package service

import (
	"context"
	system "gen/kitex_gen/system"
)

type MenuTreeService struct {
	ctx context.Context
} // NewMenuTreeService new MenuTreeService
func NewMenuTreeService(ctx context.Context) *MenuTreeService {
	return &MenuTreeService{ctx: ctx}
}

// Run create note info
func (s *MenuTreeService) Run(req *system.MenuListReq) (resp *system.MenuListResp, err error) {
	// Finish your business logic.

	return
}
