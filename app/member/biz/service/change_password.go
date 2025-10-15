package service

import (
	"context"
	Base "gen/kitex_gen/base"
	Member "gen/kitex_gen/member"
)

type ChangePasswordService struct {
	ctx context.Context
} // NewChangePasswordService new ChangePasswordService
func NewChangePasswordService(ctx context.Context) *ChangePasswordService {
	return &ChangePasswordService{ctx: ctx}
}

// Run create note info
func (s *ChangePasswordService) Run(req *Member.ChangePasswordReq) (resp *Base.BaseResp, err error) {
	return
}
