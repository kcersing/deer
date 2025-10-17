package service

import (
	"context"
	system "gen/kitex_gen/system"
)

type CreateApiService struct {
	ctx context.Context
} // NewCreateApiService new CreateApiService
func NewCreateApiService(ctx context.Context) *CreateApiService {
	return &CreateApiService{ctx: ctx}
}

// Run create note info
func (s *CreateApiService) Run(req *system.CreateApiReq) (resp *system.ApiResp, err error) {
	// Finish your business logic.

	return
}
