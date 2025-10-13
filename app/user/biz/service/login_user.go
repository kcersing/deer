package service

import (
	"context"
	User "gen/kitex_gen/user"
)

type LoginUserService struct {
	ctx context.Context
} // NewLoginUserService new LoginUserService
func NewLoginUserService(ctx context.Context) *LoginUserService {
	return &LoginUserService{ctx: ctx}
}

// Run create note info
func (s *LoginUserService) Run(req *User.CheckAccountReq) (resp *User.UserResp, err error) {
	// Finish your business logic.

	return
}
