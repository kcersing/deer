package service

import (
	"context"
	"crm/biz/convert"
	"crm/biz/dal/db"
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

	save, err := db.Client.FollowUpRecord.UpdateOneID(req.GetId()).
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
