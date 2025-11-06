package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/menu"
)

type GetMenuService struct {
	ctx context.Context
} // NewGetMenuService new MenuService
func NewGetMenuService(ctx context.Context) *GetMenuService {
	return &GetMenuService{ctx: ctx}
}

// Run create note info
func (s *GetMenuService) Run(req *base.IdReq) (resp *system.MenuResp, err error) {
	// Finish your business logic.

	only, err := db.Client.Menu.Query().Where(menu.IDEQ(req.GetId())).Only(s.ctx)
	if err != nil {
		return nil, err
	}
	resp = &system.MenuResp{
		Data: convert.EntToMenu(only),
	}

	return
}
