package service

import (
	"context"
	base "gen/kitex_gen/base"
	crm "gen/kitex_gen/crm"
)

type GetFollowUpPlanService struct {
	ctx context.Context
} // NewGetFollowUpPlanService new GetFollowUpPlanService
func NewGetFollowUpPlanService(ctx context.Context) *GetFollowUpPlanService {
	return &GetFollowUpPlanService{ctx: ctx}
}

// Run create note info
func (s *GetFollowUpPlanService) Run(req *base.IdReq) (resp *crm.FollowUpPlanResp, err error) {
	// Finish your business logic.

	return
}
