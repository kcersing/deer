package service

import (
	"context"
	"gen/kitex_gen/user"
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

	return
}
