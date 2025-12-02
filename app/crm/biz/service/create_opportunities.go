package service

import (
	"context"
	"crm/biz/convert"
	"crm/biz/dal/db"
	"gen/kitex_gen/crm"
)

type CreateOpportunitiesService struct {
	ctx context.Context
} // NewCreateOpportunitiesService new CreateOpportunitiesService
func NewCreateOpportunitiesService(ctx context.Context) *CreateOpportunitiesService {
	return &CreateOpportunitiesService{ctx: ctx}
}

// Run create note info
func (s *CreateOpportunitiesService) Run(req *crm.CreateOpportunitiesReq) (resp *crm.OpportunitiesResp, err error) {
	// Finish your business logic.

	save, err := db.Client.Opportunities.Create().
		SetTitle(req.GetTitle()).
		SetContent(req.GetContent()).
		SetUserID(req.GetUserId()).
		SetMemberID(req.GetMemberId()).
		SetPeriod(int64(req.GetPeriod())).
		SetPredictionAmount(req.GetPredictionAmount()).
		//SetCreatedID(req.GetCreatedId()).
		Save(s.ctx)

	if err != nil {
		return nil, err
	}
	resp = &crm.OpportunitiesResp{
		Data: convert.EntToOpportunities(save),
	}
	return
}
