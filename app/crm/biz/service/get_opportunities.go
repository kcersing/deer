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

type GetOpportunitiesService struct {
	ctx context.Context
} // NewGetOpportunitiesService new GetOpportunitiesService
func NewGetOpportunitiesService(ctx context.Context) *GetOpportunitiesService {
	return &GetOpportunitiesService{ctx: ctx}
}

// Run create note info
func (s *GetOpportunitiesService) Run(req *base.IdReq) (resp *crm.OpportunitiesResp, err error) {
	// Finish your business logic.

	if req == nil {
		return nil, errno.InvalidParameterErr
	}

	entity, err := db.Client.Opportunities.Get(s.ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errno.NotFound
		}
		return nil, errno.QueryFailed
	}
	dataResp := convert.EntToOpportunities(entity)
	resp = &crm.OpportunitiesResp{
		Data: dataResp,
	}
	return
}
