package service

import (
	"common/pkg/errno"
	"context"
	"crm/biz/convert"
	"crm/biz/dal/db"
	"crm/biz/dal/db/ent"
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

	if req == nil {
		return nil, errno.InvalidParameterErr
	}

	entity, err := db.Client.FollowUpPlan.Get(s.ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errno.NotFound
		}
		return nil, errno.QueryFailed
	}
	dataResp := convert.EntToFollowUpPlan(entity)
	resp = &crm.FollowUpPlanResp{
		Data: dataResp,
	}
	return
}
