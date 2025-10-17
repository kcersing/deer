package service

import (
	"context"
	base "gen/kitex_gen/base"
	"system/biz/dal/db"
)

type DeleteDicthtService struct {
	ctx context.Context
} // NewDeleteDicthtService new DeleteDicthtService
func NewDeleteDicthtService(ctx context.Context) *DeleteDicthtService {
	return &DeleteDicthtService{ctx: ctx}
}

// Run create note info
func (s *DeleteDicthtService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	err = db.Client.Dictht.DeleteOneID(req.GetId()).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
