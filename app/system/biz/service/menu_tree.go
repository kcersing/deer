package service

import (
	"context"
	"gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"github.com/cloudwego/kitex/pkg/klog"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/predicate"
)

type MenuTreeService struct {
	ctx context.Context
} // NewMenuTreeService new MenuTreeService
func NewMenuTreeService(ctx context.Context) *MenuTreeService {
	return &MenuTreeService{ctx: ctx}
}

// Run create note info
func (s *MenuTreeService) Run(req *system.MenuListReq) (resp *system.TreeResp, err error) {
	// Finish your business logic.

	var (
		predicates []predicate.Menu
		dataResp   []*base.Tree
	)
	all, err := db.Client.Menu.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return resp, err
	}
	klog.Info(all)
	dataResp = convert.FindMenuTreeChildren(all, 1)
	resp = &system.TreeResp{
		Data: dataResp,
	}
	return
}
