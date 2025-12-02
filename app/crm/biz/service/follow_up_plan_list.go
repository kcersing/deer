package service

import (
	"context"
	"crm/biz/convert"
	"crm/biz/dal/db"
	"crm/biz/dal/db/ent"
	"crm/biz/dal/db/ent/followupplan"
	"crm/biz/dal/db/ent/predicate"
	Base "gen/kitex_gen/base"
	"gen/kitex_gen/crm"
)

type FollowUpPlanListService struct {
	ctx context.Context
} // NewFollowUpPlanListService new FollowUpPlanListService
func NewFollowUpPlanListService(ctx context.Context) *FollowUpPlanListService {
	return &FollowUpPlanListService{ctx: ctx}
}

// Run create note info
func (s *FollowUpPlanListService) Run(req *crm.FollowUpPlanListReq) (resp *crm.FollowUpPlanListResp, err error) {
	// Finish your business logic.

	var (
		dataResp []*Base.FollowUpPlan
	)
	var predicates []predicate.FollowUpPlan
	if req.GetKeyword() != "" {
		predicates = append(predicates, followupplan.Or(
			followupplan.ContentContains(req.GetKeyword()),
		))
	}
	all, err := db.Client.FollowUpPlan.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(followupplan.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return nil, err
	}
	for _, e := range all {
		dataResp = append(dataResp, convert.EntToFollowUpPlan(e))
	}

	return &crm.FollowUpPlanListResp{
		Data: dataResp,
	}, nil
}
