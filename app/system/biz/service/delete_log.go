package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"system/biz/dal/db"
)

type DeleteLogService struct {
	ctx context.Context
} // NewDeleteLogService new DeleteLogService
func NewDeleteLogService(ctx context.Context) *DeleteLogService {
	return &DeleteLogService{ctx: ctx}
}

// Run create note info
func (s *DeleteLogService) Run(req *system.DeleteLogReq) (resp *base.NilResponse, err error) {
	_, err = db.Client.Logs.Delete().Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
