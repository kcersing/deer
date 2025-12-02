package service

import (
	"context"
	"crm/biz/convert"
	"crm/biz/dal/db"
	"crm/biz/dal/db/ent"
	"crm/biz/dal/db/ent/followupplan"
	"crm/biz/dal/db/ent/opportunities"
	"crm/biz/dal/db/ent/predicate"
	Base "gen/kitex_gen/base"
	crm "gen/kitex_gen/crm"
)

type OpportunitiesListService struct {
	ctx context.Context
} // NewOpportunitiesListService new OpportunitiesListService
func NewOpportunitiesListService(ctx context.Context) *OpportunitiesListService {
	return &OpportunitiesListService{ctx: ctx}
}

// Run create note info
func (s *OpportunitiesListService) Run(req *crm.OpportunitiesListReq) (resp *crm.OpportunitiesListResp, err error) {
	// Finish your business logic.

	var (
		dataResp []*Base.Opportunities
	)
	var predicates []predicate.Opportunities
	if req.GetKeyword() != "" {
		predicates = append(predicates, opportunities.Or(
			opportunities.ContentContains(req.GetKeyword()),
		))
	}
	all, err := db.Client.Opportunities.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(followupplan.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return nil, err
	}
	for _, e := range all {
		dataResp = append(dataResp, convert.EntToOpportunities(e))
	}

	return &crm.OpportunitiesListResp{
		Data: dataResp,
	}, nil
}
