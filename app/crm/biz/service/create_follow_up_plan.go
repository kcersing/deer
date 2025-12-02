package service

import (
	"common/pkg/utils"
	"context"
	"crm/biz/convert"
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
	time, err := utils.GetStringDateTime(req.GetTime())
	if err != nil {
		return nil, err
	}
	save, err := db.Client.FollowUpPlan.Create().
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
