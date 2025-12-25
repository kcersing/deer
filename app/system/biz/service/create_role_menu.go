package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"system/biz/dal/db"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
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

	err = tx.Role.UpdateOneID(req.GetRoleId()).SetMenus(req.GetIds()).Exec(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "add role's menu failed, error")
	}

	return nil, tx.Commit()
}
