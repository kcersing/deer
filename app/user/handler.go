package main

import (
	"context"
	base "gen/kitex_gen/base"
	user "gen/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *base.IdReq) (resp *user.UserResp, err error) {
	// TODO: Your code here...
	return
}
