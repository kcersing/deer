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

type CreateRoleService struct {
	ctx context.Context
} // NewCreateRoleService new CreateRoleService
func NewCreateRoleService(ctx context.Context) *CreateRoleService {
	return &CreateRoleService{ctx: ctx}
}

// Run create note info
func (s *CreateRoleService) Run(req *system.CreateRoleReq) (resp *system.RoleResp, err error) {
	// Finish your business logic.
	ok, _ := db.Client.Role.Query().Where(role.CodeEQ(req.GetCode())).Exist(s.ctx)
	if ok {
		return nil, errno.AlreadyExist
	}

	save, err := db.Client.Role.Create().
		SetName(req.GetName()).
		SetCode(req.GetCode()).
		SetDesc(req.GetDesc()).
		SetOrderNo(req.GetOrderNo()).
		SetMenus(req.GetMenus()).
		SetApis(req.GetApis()).
		SetStatus(req.GetStatus()).
		Save(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "create Role failed")
		return nil, err
	}
	resp = &system.RoleResp{
		Data: convert.EntToRole(save),
	}
	return
}
