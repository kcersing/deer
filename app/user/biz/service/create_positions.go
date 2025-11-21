package service

import (
	"context"
	"gen/kitex_gen/user"
	"user/biz/convert"
	"user/biz/dal/db"
)

type CreatePositionsService struct {
	ctx context.Context
} // NewCreatePositionsService new CreatePositionsService
func NewCreatePositionsService(ctx context.Context) *CreatePositionsService {
	return &CreatePositionsService{ctx: ctx}
}

// Run create note info
func (s *CreatePositionsService) Run(req *user.CreatePositionsReq) (resp *user.PositionsResp, err error) {
	// Finish your business logic.

	entity, err := db.Client.Position.Create().
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
