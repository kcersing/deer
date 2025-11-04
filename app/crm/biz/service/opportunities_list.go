package service

import (
	"context"
)

type OpportunitiesListService struct {
	ctx context.Context
} // NewOpportunitiesListService new OpportunitiesListService
func NewOpportunitiesListService(ctx context.Context) *OpportunitiesListService {
	return &OpportunitiesListService{ctx: ctx}
}

// Run create note info
func (s *OpportunitiesListService) Run(req *crm.OpportunitiesListReq) (resp *crm.OpportunitiesListResp, err error) {
	// Finish your business logic.

	return
}
