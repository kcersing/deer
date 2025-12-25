package service

import (
	"common/pkg/errno"
	"context"
	system "gen/kitex_gen/system"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/role"

	"github.com/pkg/errors"
)

type UpdateRoleService struct {
	ctx context.Context
} // NewUpdateRoleService new UpdateRoleService
func NewUpdateRoleService(ctx context.Context) *UpdateRoleService {
	return &UpdateRoleService{ctx: ctx}
}

// Run create note info
func (s *UpdateRoleService) Run(req *system.UpdateRoleReq) (resp *system.RoleResp, err error) {

	ok, _ := db.Client.Role.Query().Where(role.CodeEQ(req.GetCode()), role.IDNEQ(req.GetId())).Exist(s.ctx)
	if ok {
		return nil, errno.AlreadyExist
	}

	save, err := db.Client.Role.UpdateOneID(req.GetId()).
		SetName(req.GetName()).
		SetCode(req.GetCode()).
		SetDesc(req.GetDesc()).
		SetOrderNo(req.OrderNo).
		SetMenus(req.GetMenus()).
		SetApis(req.GetApis()).
		SetStatus(req.GetStatus()).
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
