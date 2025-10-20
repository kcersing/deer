package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"system/biz/dal/db"
)

type CreateRoleMenuService struct {
	ctx context.Context
} // NewCreateRoleMenuService new CreateRoleMenuService
func NewCreateRoleMenuService(ctx context.Context) *CreateRoleMenuService {
	return &CreateRoleMenuService{ctx: ctx}
}

// Run create note info
func (s *CreateRoleMenuService) Run(req *system.CreateMenuAuthReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	tx, err := db.Client.Tx(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "starting a transaction err")
	}
	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				hlog.Error("UpdateMenuAuthority err:", err, "rollback err:", rollbackErr)
			}
		}
	}()

	//tx.Role.UpdateOneID(roleID).ClearMenus().Exec(a.ctx)
	err = tx.Role.UpdateOneID(req.GetRoleId()).ClearMenus().Exec(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "delete role's menu failed, error")
	}

	err = tx.Role.UpdateOneID(req.GetRoleId()).AddMenuIDs(req.GetIds()...).Exec(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "add role's menu failed, error")
	}

	return nil, tx.Commit()
}
