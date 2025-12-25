package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent"
	"system/biz/dal/db/ent/menu"
	"system/biz/dal/db/ent/role"
)

type GetRoleMenuService struct {
	ctx context.Context
} // NewGetRoleMenuService new GetRoleMenuService
func NewGetRoleMenuService(ctx context.Context) *GetRoleMenuService {
	return &GetRoleMenuService{ctx: ctx}
}

// Run create note info
func (s *GetRoleMenuService) Run(req *base.IdReq) (resp *system.MenuListResp, err error) {
	// Finish your business logic.
	f, err := db.Client.Role.
		Query().
		Where(role.IDIn(req.GetId())).First(s.ctx)
	if err != nil {
		return nil, err
	}
	menus, err := db.Client.Menu.
		Query().
		Where(menu.IDIn(f.Menus...)).
		Order(ent.Asc(menu.FieldID)).
		All(s.ctx)
	if err != nil {
		return nil, err
	}
	resp = &system.MenuListResp{
		Data: convert.FindMenuChildren(menus, 1),
	}
	return
}
