package service

import (
	"context"
	"gen/kitex_gen/user"
)

type CreatePositionsService struct {
	ctx context.Context
} // NewCreatePositionsService new CreatePositionsService
func NewCreatePositionsService(ctx context.Context) *CreatePositionsService {
	return &CreatePositionsService{ctx: ctx}
}

// Run create note info
func (s *CreatePositionsService) Run(req *user.CreatePositionsReq) (resp *user.PositionsResp, err error) {
	// Finish your business logic.

	return
}
