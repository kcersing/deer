package service

import (
	"context"
	base "gen/kitex_gen/base"
)

type DeletePropertyService struct {
	ctx context.Context
} // NewDeletePropertyService new DeletePropertyService
func NewDeletePropertyService(ctx context.Context) *DeletePropertyService {
	return &DeletePropertyService{ctx: ctx}
}

// Run create note info
func (s *DeletePropertyService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
