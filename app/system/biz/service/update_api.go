package service

import (
	"context"
	system "gen/kitex_gen/system"
)

type UpdateApiService struct {
	ctx context.Context
} // NewUpdateApiService new UpdateApiService
func NewUpdateApiService(ctx context.Context) *UpdateApiService {
	return &UpdateApiService{ctx: ctx}
}

// Run create note info
func (s *UpdateApiService) Run(req *system.UpdateApiReq) (resp *system.ApiResp, err error) {
	// Finish your business logic.

	return
}
