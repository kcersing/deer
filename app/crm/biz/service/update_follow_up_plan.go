package service

import (
	"common/pkg/utils"
	"context"
	"crm/biz/convert"
	"crm/biz/dal/db"
	crm "gen/kitex_gen/crm"
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

	time, err := utils.GetStringDateTime(req.GetTime())
	if err != nil {
		return nil, err
	}
	save, err := db.Client.FollowUpPlan.UpdateOneID(req.GetId()).
		SetContent(req.GetContent()).
		SetTime(time).
		SetMemberID(req.GetMemberId()).
		SetUserID(req.GetUserId()).
		SetStatus(req.GetStatus()).
		SetCreatedID(req.GetCreatedId()).
		SetDivision(req.GetDivision()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	resp = &crm.FollowUpPlanResp{
		Data: convert.EntToFollowUpPlan(save),
	}
	return
}
