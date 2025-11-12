package service

import (
	"context"
	"crm/biz/dal/db"
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
	db.Client.FollowUpPlan.Create().
		SetContent(req.GetContent()).
		SetTime(req.GetTime()).
		SetMemberID(req.GetMemberId()).
		SetUserID(req.GetUserId()).
		SetStatus(req.GetStatus()).
		SetCreatedID(req.GetCreatedId()).
		SetDivision(req.GetDivision()).
		Save(s.ctx)
	return
}
