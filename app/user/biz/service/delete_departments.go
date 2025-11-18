package service

import (
	"context"
	"gen/kitex_gen/base"
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

	return
}
