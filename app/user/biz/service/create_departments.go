package service

import (
	"context"
	"gen/kitex_gen/user"
	"user/biz/convert"
	"user/biz/dal/db"
)

type CreateDepartmentsService struct {
	ctx context.Context
} // NewCreateDepartmentsService new CreateDepartmentsService
func NewCreateDepartmentsService(ctx context.Context) *CreateDepartmentsService {
	return &CreateDepartmentsService{ctx: ctx}
}

// Run create note info
func (s *CreateDepartmentsService) Run(req *user.CreateDepartmentsReq) (resp *user.DepartmentsResp, err error) {
	// Finish your business logic.

	entity, err := db.Client.Department.Create().
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
