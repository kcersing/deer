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

type GetFollowUpRecordService struct {
	ctx context.Context
} // NewGetFollowUpRecordService new GetFollowUpRecordService
func NewGetFollowUpRecordService(ctx context.Context) *GetFollowUpRecordService {
	return &GetFollowUpRecordService{ctx: ctx}
}

// Run create note info
func (s *GetFollowUpRecordService) Run(req *base.IdReq) (resp *crm.FollowUpRecordResp, err error) {
	// Finish your business logic.

	if req == nil {
		return nil, errno.InvalidParameterErr
	}

	entity, err := db.Client.FollowUpRecord.Get(s.ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errno.NotFound
		}
		return nil, errno.QueryFailed
	}
	dataResp := convert.EntToFollowUpRecord(entity)
	resp = &crm.FollowUpRecordResp{
		Data: dataResp,
	}
	return
}
