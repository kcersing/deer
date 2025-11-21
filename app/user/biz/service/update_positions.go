package service

import (
	"context"
	"gen/kitex_gen/user"
	"user/biz/convert"
	"user/biz/dal/db"
)

type UpdatePositionsService struct {
	ctx context.Context
} // NewUpdatePositionsService new UpdatePositionsService
func NewUpdatePositionsService(ctx context.Context) *UpdatePositionsService {
	return &UpdatePositionsService{ctx: ctx}
}

// Run create note info
func (s *UpdatePositionsService) Run(req *user.UpdatePositionsReq) (resp *user.PositionsResp, err error) {
	// Finish your business logic.

	entity, err := db.Client.Position.UpdateOneID(req.GetId()).
		SetName(req.GetName()).
		SetCode(req.GetCode()).
		SetDepartmentID(req.GetDepartmentId()).
		SetParentID(req.GetParentId()).
		SetDesc(req.GetDesc()).
		SetStatus(req.GetStatus()).
		SetQuota(req.GetQuota()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	dataResp := convert.EntToPosition(entity)
	resp = &user.PositionsResp{
		Data: dataResp,
	}
	return
}
