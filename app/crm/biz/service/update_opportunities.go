package service

import (
	"context"
	"crm/biz/convert"
	"crm/biz/dal/db"
	crm "gen/kitex_gen/crm"
)

type UpdateOpportunitiesService struct {
	ctx context.Context
} // NewUpdateOpportunitiesService new UpdateOpportunitiesService
func NewUpdateOpportunitiesService(ctx context.Context) *UpdateOpportunitiesService {
	return &UpdateOpportunitiesService{ctx: ctx}
}

// Run create note info
func (s *UpdateOpportunitiesService) Run(req *crm.UpdateOpportunitiesReq) (resp *crm.OpportunitiesResp, err error) {
	// Finish your business logic.

	save, err := db.Client.Opportunities.UpdateOneID(req.GetId()).
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
