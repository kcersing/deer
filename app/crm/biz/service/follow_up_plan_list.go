package service

import (
	"context"
)

type FollowUpPlanListService struct {
	ctx context.Context
} // NewFollowUpPlanListService new FollowUpPlanListService
func NewFollowUpPlanListService(ctx context.Context) *FollowUpPlanListService {
	return &FollowUpPlanListService{ctx: ctx}
}

// Run create note info
func (s *FollowUpPlanListService) Run(req *crm.FollowUpPlanListReq) (resp *crm.FollowUpPlanListResp, err error) {
	// Finish your business logic.

	return
}
