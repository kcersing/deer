package service

import (
	"context"
	"crm/biz/convert"
	"crm/biz/dal/db"
	"gen/kitex_gen/crm"
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

	save, err := db.Client.FollowUpRecord.Create().
		SetFollowUpID(req.GetFollowUpId()).
		SetContent(req.GetContent()).
		SetMethod(req.GetMethod()).
		SetUserID(req.GetUserId()).
		SetDivision(req.GetDivision()).
		SetRecord(req.GetRecord()).
		SetOpportunitiesID(req.GetOpportunitiesId()).
		//SetCreatedID(req.GetCreatedId()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	resp = &crm.FollowUpRecordResp{
		Data: convert.EntToFollowUpRecord(save),
	}
	return
}
