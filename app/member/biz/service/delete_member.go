package service

import (
	"context"
	Base "gen/kitex_gen/base"
)

type DeleteMemberService struct {
	ctx context.Context
} // NewDeleteMemberService new DeleteMemberService
func NewDeleteMemberService(ctx context.Context) *DeleteMemberService {
	return &DeleteMemberService{ctx: ctx}
}

// Run create note info
func (s *DeleteMemberService) Run(req *Base.IdReq) (resp *Base.BaseResp, err error) {
	return
}
