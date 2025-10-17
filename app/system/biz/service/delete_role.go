package service

import (
	"context"
	base "gen/kitex_gen/base"
	"system/biz/dal/db"
)

type DeleteRoleService struct {
	ctx context.Context
} // NewDeleteRoleService new DeleteRoleService
func NewDeleteRoleService(ctx context.Context) *DeleteRoleService {
	return &DeleteRoleService{ctx: ctx}
}

// Run create note info
func (s *DeleteRoleService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	err = db.Client.Role.DeleteOneID(req.GetId()).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
