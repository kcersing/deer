package service

import (
	"context"
	system "gen/kitex_gen/system"
	"github.com/pkg/errors"
	"system/biz/convert"
	"system/biz/dal/db"
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
	save, err := db.Client.Role.Create().
		SetName(req.Name).
		SetValue(req.Value).
		SetDefaultRouter(req.DefaultRouter).
		SetRemark(req.Remark).
		//SetOrderNo(req.OrderNo).

		Save(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "create Role failed")
		return nil, err
	}
	resp.Data = convert.EntToRole(save)
	return
}
