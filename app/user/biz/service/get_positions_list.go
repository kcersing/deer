package service

import (
	"context"
	"gen/kitex_gen/user"
)

type GetPositionsListService struct {
	ctx context.Context
} // NewGetPositionsListService new GetPositionsListService
func NewGetPositionsListService(ctx context.Context) *GetPositionsListService {
	return &GetPositionsListService{ctx: ctx}
}

// Run create note info
func (s *GetPositionsListService) Run(req *user.GetPositionsListReq) (resp *user.PositionsListResp, err error) {
	// Finish your business logic.

	return
}
