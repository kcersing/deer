package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
)

type DeleteLogService struct {
	ctx context.Context
} // NewDeleteLogService new DeleteLogService
func NewDeleteLogService(ctx context.Context) *DeleteLogService {
	return &DeleteLogService{ctx: ctx}
}

// Run create note info
func (s *DeleteLogService) Run(req *system.DeleteLog) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
