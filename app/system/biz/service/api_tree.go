package service

import (
	"context"
	system "gen/kitex_gen/system"
)

type ApiTreeService struct {
	ctx context.Context
} // NewApiTreeService new ApiTreeService
func NewApiTreeService(ctx context.Context) *ApiTreeService {
	return &ApiTreeService{ctx: ctx}
}

// Run create note info
func (s *ApiTreeService) Run(req *system.ApiListReq) (resp *system.ApiListResp, err error) {
	// Finish your business logic.

	return
}
