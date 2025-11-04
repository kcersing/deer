package service

import (
	"context"
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

	return
}
