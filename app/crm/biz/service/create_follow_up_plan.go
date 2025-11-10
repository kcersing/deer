package service

import (
	"context"
	"gen/kitex_gen/crm"
)

type CreateFollowUpPlanService struct {
	ctx context.Context
} // NewCreateFollowUpPlanService new CreateFollowUpPlanService
func NewCreateFollowUpPlanService(ctx context.Context) *CreateFollowUpPlanService {
	return &CreateFollowUpPlanService{ctx: ctx}
}

// Run create note info
func (s *CreateFollowUpPlanService) Run(req *crm.CreateFollowUpPlanReq) (resp *crm.FollowUpPlanResp, err error) {
	// Finish your business logic.

	return
}
