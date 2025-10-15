package service

import (
	"context"
	Base "gen/kitex_gen/base"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *Base.IdReq) (resp *Base.BaseResp, err error) {
	return
}
