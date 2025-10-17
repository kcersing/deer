package service

import (
	"context"
	base "gen/kitex_gen/base"
)

type DeleteDicthtService struct {
	ctx context.Context
} // NewDeleteDicthtService new DeleteDicthtService
func NewDeleteDicthtService(ctx context.Context) *DeleteDicthtService {
	return &DeleteDicthtService{ctx: ctx}
}

// Run create note info
func (s *DeleteDicthtService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
