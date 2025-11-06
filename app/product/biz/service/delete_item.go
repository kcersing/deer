package service

import (
	"context"
	base "gen/kitex_gen/base"
)

type DeleteItemService struct {
	ctx context.Context
} // NewDeleteItemService new DeleteItemService
func NewDeleteItemService(ctx context.Context) *DeleteItemService {
	return &DeleteItemService{ctx: ctx}
}

// Run create note info
func (s *DeleteItemService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
