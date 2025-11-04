package service

import (
	"context"
)

type CreateFollowUpRecordService struct {
	ctx context.Context
} // NewCreateFollowUpRecordService new CreateFollowUpRecordService
func NewCreateFollowUpRecordService(ctx context.Context) *CreateFollowUpRecordService {
	return &CreateFollowUpRecordService{ctx: ctx}
}

// Run create note info
func (s *CreateFollowUpRecordService) Run(req *crm.CreateFollowUpRecordReq) (resp *crm.FollowUpRecordResp, err error) {
	// Finish your business logic.

	return
}
