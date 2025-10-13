package service

import (
	"context"
	User "gen/kitex_gen/user"
)

type CreateUserService struct {
	ctx context.Context
} // NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// Run create note info
func (s *CreateUserService) Run(req *User.CreateUserReq) (resp *User.UserResp, err error) {
	// Finish your business logic.

	return
}
