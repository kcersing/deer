package service

import (
	"context"
	"gen/kitex_gen/base"
	"user/biz/dal/db"
	"user/biz/dal/db/ent/department"
)

type DeleteDepartmentsService struct {
	ctx context.Context
} // NewDeleteDepartmentsService new DeleteDepartmentsService
func NewDeleteDepartmentsService(ctx context.Context) *DeleteDepartmentsService {
	return &DeleteDepartmentsService{ctx: ctx}
}

// Run create note info
func (s *DeleteDepartmentsService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	_, err = db.Client.Department.Delete().Where(department.IDEQ(req.GetId())).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
