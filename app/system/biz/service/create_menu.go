package service

import (
	"context"
	system "gen/kitex_gen/system"
	"github.com/pkg/errors"
	"system/biz/convert"
	"system/biz/dal/db"
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
	save, err := db.Client.Menu.Create().
		SetParentID(req.GetParentId()).
		SetPath(req.GetPath()).
		SetName(req.GetName()).
		//SetOrderNo(req.GetOrderNo()).
		//SetDisabled(req.GetDisabled()).
		//SetIgnore(req.GetIgnore()).
		SetType(req.GetType()).
		SetLevel(req.GetLevel()).
		SetMenuType(req.GetMenuType()).
		SetRedirect(req.GetRedirect()).
		SetComponent(req.GetComponent()).
		SetTitle(req.GetTitle()).
		//SetIcon(req.Icon).
		SetHidden(req.GetHidden()).
		SetSort(req.GetSort()).
		SetURL(req.GetUrl()).
		SetStatus(req.GetStatus()).
		Save(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "create menu failed")
	}
	resp.Data = convert.EntToMenu(save)
	return
}
