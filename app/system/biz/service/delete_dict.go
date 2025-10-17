package service

import (
	"context"
	base "gen/kitex_gen/base"
)

type DeleteDictService struct {
	ctx context.Context
} // NewDeleteDictService new DeleteDictService
func NewDeleteDictService(ctx context.Context) *DeleteDictService {
	return &DeleteDictService{ctx: ctx}
}

// Run create note info
func (s *DeleteDictService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
