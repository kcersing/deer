package service

import (
	"context"
	system "gen/kitex_gen/system"
	"github.com/pkg/errors"
	"system/biz/convert"
	"system/biz/dal/db"
)

type UpdateRoleService struct {
	ctx context.Context
} // NewUpdateRoleService new UpdateRoleService
func NewUpdateRoleService(ctx context.Context) *UpdateRoleService {
	return &UpdateRoleService{ctx: ctx}
}

// Run create note info
func (s *UpdateRoleService) Run(req *system.UpdateRoleReq) (resp *system.RoleResp, err error) {
	// Finish your business logic.

	save, err := db.Client.Role.UpdateOneID(req.GetId()).
		SetName(req.GetName()).
		SetCode(req.GetCode()).
		SetDesc(req.GetDesc()).
		SetOrderNo(req.OrderNo).
		Save(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Role failed")
		return nil, err
	}
	resp = &system.RoleResp{
		Data: convert.EntToRole(save),
	}
	return
}
