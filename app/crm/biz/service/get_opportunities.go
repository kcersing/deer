package service

import (
	"context"
	base "gen/kitex_gen/base"
	crm "gen/kitex_gen/crm"
)

type GetOpportunitiesService struct {
	ctx context.Context
} // NewGetOpportunitiesService new GetOpportunitiesService
func NewGetOpportunitiesService(ctx context.Context) *GetOpportunitiesService {
	return &GetOpportunitiesService{ctx: ctx}
}

// Run create note info
func (s *GetOpportunitiesService) Run(req *base.IdReq) (resp *crm.OpportunitiesResp, err error) {
	// Finish your business logic.

	return
}
