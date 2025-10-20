package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"system/biz/dal/db"
)

type CreateRoleApiService struct {
	ctx context.Context
} // NewCreateRoleApiService new CreateRoleApiService
func NewCreateRoleApiService(ctx context.Context) *CreateRoleApiService {
	return &CreateRoleApiService{ctx: ctx}
}

// Run create note info
func (s *CreateRoleApiService) Run(req *system.CreateMenuAuthReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	// Finish your business logic.
	tx, err := db.Client.Tx(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "starting a transaction err")
	}
	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				hlog.Error("UpdateAPIAuthority err:", err, "rollback err:", rollbackErr)
			}
		}
	}()

	//tx.Role.UpdateOneID(roleID).ClearAPI().Exec(a.ctx)
	err = tx.Role.UpdateOneID(req.GetRoleId()).ClearAPI().Exec(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "delete role's api failed, error")
	}

	err = tx.Role.UpdateOneID(req.GetRoleId()).AddAPIIDs(req.GetIds()...).Exec(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "add role's api failed, error")
	}

	return nil, tx.Commit()
}
