package service

import (
	"context"
	base "gen/kitex_gen/base"
)

type DeleteOpportunitiesService struct {
	ctx context.Context
} // NewDeleteOpportunitiesService new DeleteOpportunitiesService
func NewDeleteOpportunitiesService(ctx context.Context) *DeleteOpportunitiesService {
	return &DeleteOpportunitiesService{ctx: ctx}
}

// Run create note info
func (s *DeleteOpportunitiesService) Run(req *base.IdReq) (resp *base.BaseResp, err error) {
	// Finish your business logic.

	return
}
