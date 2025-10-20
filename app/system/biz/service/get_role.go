package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/role"
)

type GetRoleService struct {
	ctx context.Context
} // NewGetRoleService new GetRoleService
func NewGetRoleService(ctx context.Context) *GetRoleService {
	return &GetRoleService{ctx: ctx}
}

// Run create note info
func (s *GetRoleService) Run(req *base.IdReq) (resp *system.RoleResp, err error) {
	// Finish your business logic.
	only, err := db.Client.Role.Query().Where(role.IDEQ(req.GetId())).Only(s.ctx)
	if err != nil {
		return nil, err
	}
	resp.Data = convert.EntToRole(only)
	return
}
