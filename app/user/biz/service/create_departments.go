package service

import (
	"context"
	"gen/kitex_gen/user"
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

	return
}
