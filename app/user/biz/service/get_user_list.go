package service

import (
	"context"
	User "gen/kitex_gen/user"
)

type GetUserListService struct {
	ctx context.Context
} // NewGetUserListService new GetUserListService
func NewGetUserListService(ctx context.Context) *GetUserListService {
	return &GetUserListService{ctx: ctx}
}

// Run Update
func (s *GetUserListService) Run(req *User.GetUserListReq) (resp *User.UserResp, err error) {

	return
}
