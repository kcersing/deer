package service

import (
	"context"
	base "gen/kitex_gen/base"
	"system/biz/dal/db"
)

type DeleteApiService struct {
	ctx context.Context
} // NewDeleteApiService new DeleteApiService
func NewDeleteApiService(ctx context.Context) *DeleteApiService {
	return &DeleteApiService{ctx: ctx}
}

// Run create note info
func (s *DeleteApiService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	err = db.Client.API.DeleteOneID(req.GetId()).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
