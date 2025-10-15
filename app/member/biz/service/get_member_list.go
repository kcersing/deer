package service

import (
	"context"
	Member "gen/kitex_gen/member"
)

type GetMemberListService struct {
	ctx context.Context
} // NewGetMemberListService new GetMemberListService
func NewGetMemberListService(ctx context.Context) *GetMemberListService {
	return &GetMemberListService{ctx: ctx}
}

// Run Update
func (s *GetMemberListService) Run(req *Member.GetMemberListReq) (resp *Member.MemberListResp, err error) {

	return
}
