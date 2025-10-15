package service

import (
	"context"
	Base "gen/kitex_gen/base"
	User "gen/kitex_gen/user"
)

type ChangePasswordService struct {
	ctx context.Context
} // NewChangePasswordService new ChangePasswordService
func NewChangePasswordService(ctx context.Context) *ChangePasswordService {
	return &ChangePasswordService{ctx: ctx}
}

// Run create note info
func (s *ChangePasswordService) Run(req *User.ChangePasswordReq) (resp *Base.BaseResp, err error) {
	return
}
