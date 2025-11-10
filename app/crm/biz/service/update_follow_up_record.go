package service

import (
	"context"
	crm "gen/kitex_gen/crm"
)

type UpdateFollowUpRecordService struct {
	ctx context.Context
} // NewUpdateFollowUpRecordService new UpdateFollowUpRecordService
func NewUpdateFollowUpRecordService(ctx context.Context) *UpdateFollowUpRecordService {
	return &UpdateFollowUpRecordService{ctx: ctx}
}

// Run create note info
func (s *UpdateFollowUpRecordService) Run(req *crm.UpdateFollowUpRecordReq) (resp *crm.FollowUpRecordResp, err error) {
	// Finish your business logic.

	return
}
