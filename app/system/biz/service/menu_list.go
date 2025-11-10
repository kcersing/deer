package service

import (
	"context"
	"gen/kitex_gen/base"
	"gen/kitex_gen/system"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/predicate"
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
	var (
		predicates []predicate.Menu
		dataResp   []*base.Menu
	)
	all, err := db.Client.Menu.Query().Where(predicates...).
		All(s.ctx)
	if err != nil {
		return resp, err
	}

	dataResp = convert.FindMenuChildren(all, 1)
	resp = &system.MenuListResp{
		Data: dataResp,
	}
	return
}
