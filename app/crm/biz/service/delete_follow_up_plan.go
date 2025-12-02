package service

import (
	"context"
	"crm/biz/dal/db"
	"crm/biz/dal/db/ent/followupplan"
	base "gen/kitex_gen/base"
)

type DeleteFollowUpPlanService struct {
	ctx context.Context
} // NewDeleteFollowUpPlanService new DeleteFollowUpPlanService
func NewDeleteFollowUpPlanService(ctx context.Context) *DeleteFollowUpPlanService {
	return &DeleteFollowUpPlanService{ctx: ctx}
}

// Run create note info
func (s *DeleteFollowUpPlanService) Run(req *base.IdReq) (resp *base.BaseResp, err error) {
	// Finish your business logic.
	_, err = db.Client.FollowUpPlan.Delete().Where(followupplan.IDEQ(req.GetId())).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
