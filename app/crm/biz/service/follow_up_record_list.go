package service

import (
	"context"
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

	return
}
