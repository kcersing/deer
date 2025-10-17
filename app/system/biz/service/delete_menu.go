package service

import (
	"context"
	base "gen/kitex_gen/base"
	"system/biz/dal/db"
)

type DeleteMenuService struct {
	ctx context.Context
} // NewDeleteMenuService new DeleteMenuService
func NewDeleteMenuService(ctx context.Context) *DeleteMenuService {
	return &DeleteMenuService{ctx: ctx}
}

// Run create note info
func (s *DeleteMenuService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	err = db.Client.Menu.DeleteOneID(req.GetId()).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
