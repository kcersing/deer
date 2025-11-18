package service

import (
	"common/pkg/errno"
	"context"
	"gen/kitex_gen/base"
	"gen/kitex_gen/user"
	"user/biz/convert"
	"user/biz/dal/db"
	"user/biz/dal/db/ent"
)

type GetDepartmentsService struct {
	ctx context.Context
} // NewGetDepartmentsService new GetDepartmentsService
func NewGetDepartmentsService(ctx context.Context) *GetDepartmentsService {
	return &GetDepartmentsService{ctx: ctx}
}

// Run create note info
func (s *GetDepartmentsService) Run(req *base.IdReq) (resp *user.DepartmentsResp, err error) {

	if req == nil {
		return nil, errno.InvalidParameterErr
	}

	entity, err := db.Client.Department.Get(s.ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errno.NotFound
		}
		return nil, errno.QueryFailed
	}
	dataResp := convert.EntToDepartments(entity)
	resp = &user.DepartmentsResp{
		Data: dataResp,
	}
	return
}
