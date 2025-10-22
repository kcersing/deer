package main

import (
	"context"
	base "gen/kitex_gen/base"
	user "gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"

	"user/biz/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserReq) (resp *user.UserResp, err error) {
	resp, err = service.NewCreateUserService(ctx).Run(req)

	return resp, err
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *base.IdReq) (resp *user.UserResp, err error) {
	resp, err = service.NewGetUserService(ctx).Run(req)

	return resp, err
}

// GetUserList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserList(ctx context.Context, req *user.GetUserListReq) (resp *user.UserListResp, err error) {
	resp, err = service.NewGetUserListService(ctx).Run(req)

	return resp, err
}

// LoginUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginUser(ctx context.Context, req *base.CheckAccountReq) (resp *user.UserResp, err error) {

	klog.Error("111111111111111111111111111111111111111")
	//resp, err = service.NewLoginUserService(ctx).Run(req)
	return resp, err
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserReq) (resp *user.UserResp, err error) {

	resp, err = service.NewUpdateUserService(ctx).Run(req)

	return resp, err
}

// ChangePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangePassword(ctx context.Context, req *user.ChangePasswordReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewChangePasswordService(ctx).Run(req)

	return resp, err
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteUserService(ctx).Run(req)

	return resp, err
}

// SetUserRole implements the UserServiceImpl interface.
func (s *UserServiceImpl) SetUserRole(ctx context.Context, req *user.SetUserRoleReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewSetUserRoleService(ctx).Run(req)

	return resp, err
}
