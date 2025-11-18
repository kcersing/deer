package service

import (
	"context"
	"gen/kitex_gen/user"
)

type UpdatePositionsService struct {
	ctx context.Context
} // NewUpdatePositionsService new UpdatePositionsService
func NewUpdatePositionsService(ctx context.Context) *UpdatePositionsService {
	return &UpdatePositionsService{ctx: ctx}
}

// Run create note info
func (s *UpdatePositionsService) Run(req *user.UpdatePositionsReq) (resp *user.PositionsResp, err error) {
	// Finish your business logic.

	return
}
