package service

import (
	"context"
	"crm/biz/convert"
	"crm/biz/dal/db"
	"crm/biz/dal/db/ent"
	"crm/biz/dal/db/ent/followupplan"
	"crm/biz/dal/db/ent/followuprecord"
	"crm/biz/dal/db/ent/predicate"
	Base "gen/kitex_gen/base"
	"gen/kitex_gen/crm"
)

type FollowUpRecordListService struct {
	ctx context.Context
} // NewFollowUpRecordListService new FollowUpRecordListService
func NewFollowUpRecordListService(ctx context.Context) *FollowUpRecordListService {
	return &FollowUpRecordListService{ctx: ctx}
}

// Run create note info
func (s *FollowUpRecordListService) Run(req *crm.FollowUpRecordListReq) (resp *crm.FollowUpRecordListResp, err error) {
	// Finish your business logic.

	var (
		dataResp []*Base.FollowUpRecord
	)
	var predicates []predicate.FollowUpRecord
	if req.GetKeyword() != "" {
		predicates = append(predicates, followuprecord.Or(
			followuprecord.ContentContains(req.GetKeyword()),
		))
	}
	all, err := db.Client.FollowUpRecord.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(followupplan.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return nil, err
	}
	for _, e := range all {
		dataResp = append(dataResp, convert.EntToFollowUpRecord(e))
	}

	return &crm.FollowUpRecordListResp{
		Data: dataResp,
	}, nil
}
