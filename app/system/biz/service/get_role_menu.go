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
	menus, err := db.Client.Role.
		Query().
		Where(role.IDIn(req.GetId())).
		QueryMenu().
		Where(menu.DeleteEQ(0)).
		//WithChildren().
		Order(ent.Asc(menu.FieldOrderNo)).
		All(s.ctx)
	if err != nil {
		return nil, err
	}
	resp = &system.MenuListResp{
		Data: convert.FindMenuChildren(menus, 1),
	}
	return
}
