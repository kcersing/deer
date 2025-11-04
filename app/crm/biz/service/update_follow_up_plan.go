package service

import (
	"context"
)

type UpdateFollowUpPlanService struct {
	ctx context.Context
} // NewUpdateFollowUpPlanService new UpdateFollowUpPlanService
func NewUpdateFollowUpPlanService(ctx context.Context) *UpdateFollowUpPlanService {
	return &UpdateFollowUpPlanService{ctx: ctx}
}

// Run create note info
func (s *UpdateFollowUpPlanService) Run(req *crm.UpdateFollowUpPlanReq) (resp *crm.FollowUpPlanResp, err error) {
	// Finish your business logic.

	return
}
