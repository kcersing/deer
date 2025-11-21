package service

import (
	"context"
	"gen/kitex_gen/user"
	"user/biz/convert"
	"user/biz/dal/db"
)

type UpdateDepartmentsService struct {
	ctx context.Context
} // NewUpdateDepartmentsService new UpdateDepartmentsService
func NewUpdateDepartmentsService(ctx context.Context) *UpdateDepartmentsService {
	return &UpdateDepartmentsService{ctx: ctx}
}

// Run create note info
func (s *UpdateDepartmentsService) Run(req *user.UpdateDepartmentsReq) (resp *user.DepartmentsResp, err error) {
	// Finish your business logic.

	entity, err := db.Client.Department.UpdateOneID(req.GetId()).
		SetName(req.GetName()).
		SetManagerID(req.GetManagerId()).
		SetParentID(req.GetParentId()).
		SetDesc(req.GetDesc()).
		SetStatus(req.GetStatus()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	dataResp := convert.EntToDepartments(entity)
	resp = &user.DepartmentsResp{
		Data: dataResp,
	}
	return
}
