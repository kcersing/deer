package service

import (
	"context"
	"gen/kitex_gen/base"
)

type DeletePositionsService struct {
	ctx context.Context
} // NewDeletePositionsService new DeletePositionsService
func NewDeletePositionsService(ctx context.Context) *DeletePositionsService {
	return &DeletePositionsService{ctx: ctx}
}

// Run create note info
func (s *DeletePositionsService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
