package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
)

type MenuService struct {
	ctx context.Context
} // NewMenuService new MenuService
func NewMenuService(ctx context.Context) *MenuService {
	return &MenuService{ctx: ctx}
}

// Run create note info
func (s *MenuService) Run(req *base.IdReq) (resp *system.MenuResp, err error) {
	// Finish your business logic.

	return
}
