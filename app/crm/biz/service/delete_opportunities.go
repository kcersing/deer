package service

import (
	"context"
	"crm/biz/dal/db"
	"crm/biz/dal/db/ent/opportunities"
	base "gen/kitex_gen/base"
)

type DeleteOpportunitiesService struct {
	ctx context.Context
} // NewDeleteOpportunitiesService new DeleteOpportunitiesService
func NewDeleteOpportunitiesService(ctx context.Context) *DeleteOpportunitiesService {
	return &DeleteOpportunitiesService{ctx: ctx}
}

// Run create note info
func (s *DeleteOpportunitiesService) Run(req *base.IdReq) (resp *base.BaseResp, err error) {
	// Finish your business logic.

	_, err = db.Client.Opportunities.Delete().Where(opportunities.IDEQ(req.GetId())).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
