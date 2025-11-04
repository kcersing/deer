package service

import (
	"context"
	base "gen/kitex_gen/base"
)

type DeleteFollowUpPlanService struct {
	ctx context.Context
} // NewDeleteFollowUpPlanService new DeleteFollowUpPlanService
func NewDeleteFollowUpPlanService(ctx context.Context) *DeleteFollowUpPlanService {
	return &DeleteFollowUpPlanService{ctx: ctx}
}

// Run create note info
func (s *DeleteFollowUpPlanService) Run(req *base.IdReq) (resp *base.BaseResp, err error) {
	// Finish your business logic.

	return
}
