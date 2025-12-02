package service

import (
	"context"
	"crm/biz/dal/db"
	"crm/biz/dal/db/ent/followuprecord"
	base "gen/kitex_gen/base"
)

type DeleteFollowUpRecordService struct {
	ctx context.Context
} // NewDeleteFollowUpRecordService new DeleteFollowUpRecordService
func NewDeleteFollowUpRecordService(ctx context.Context) *DeleteFollowUpRecordService {
	return &DeleteFollowUpRecordService{ctx: ctx}
}

// Run create note info
func (s *DeleteFollowUpRecordService) Run(req *base.IdReq) (resp *base.BaseResp, err error) {
	// Finish your business logic.

	_, err = db.Client.FollowUpRecord.Delete().Where(followuprecord.IDEQ(req.GetId())).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
