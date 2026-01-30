package main

import (
	"context"
	base "gen/kitex_gen/base"
	user "gen/kitex_gen/user"
	"user/biz/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
}

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

	resp, err = service.NewLoginUserService(ctx).Run(req)
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

// CreateDepartments implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateDepartments(ctx context.Context, req *user.CreateDepartmentsReq) (resp *user.DepartmentsResp, err error) {
	resp, err = service.NewCreateDepartmentsService(ctx).Run(req)

	return resp, err
}

// DeleteDepartments implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteDepartments(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteDepartmentsService(ctx).Run(req)

	return resp, err
}

// UpdateDepartments implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateDepartments(ctx context.Context, req *user.UpdateDepartmentsReq) (resp *user.DepartmentsResp, err error) {
	resp, err = service.NewUpdateDepartmentsService(ctx).Run(req)

	return resp, err
}

// GetDepartments implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetDepartments(ctx context.Context, req *base.IdReq) (resp *user.DepartmentsResp, err error) {
	resp, err = service.NewGetDepartmentsService(ctx).Run(req)

	return resp, err
}

// GetDepartmentsList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetDepartmentsList(ctx context.Context, req *user.GetDepartmentsListReq) (resp *user.DepartmentsListResp, err error) {
	resp, err = service.NewGetDepartmentsListService(ctx).Run(req)

	return resp, err
}

// CreatePositions implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreatePositions(ctx context.Context, req *user.CreatePositionsReq) (resp *user.PositionsResp, err error) {
	resp, err = service.NewCreatePositionsService(ctx).Run(req)

	return resp, err
}

// DeletePositions implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeletePositions(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeletePositionsService(ctx).Run(req)

	return resp, err
}

// UpdatePositions implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdatePositions(ctx context.Context, req *user.UpdatePositionsReq) (resp *user.PositionsResp, err error) {
	resp, err = service.NewUpdatePositionsService(ctx).Run(req)

	return resp, err
}

// GetPositions implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetPositions(ctx context.Context, req *base.IdReq) (resp *user.PositionsResp, err error) {
	resp, err = service.NewGetPositionsService(ctx).Run(req)

	return resp, err
}

// GetPositionsList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetPositionsList(ctx context.Context, req *user.GetPositionsListReq) (resp *user.PositionsListResp, err error) {
	resp, err = service.NewGetPositionsListService(ctx).Run(req)

	return resp, err
}

// GetUserIds implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserIds(ctx context.Context, req *user.GetUserListReq) (resp *user.UserIdsResp, err error) {
	// TODO: Your code here...
	return
}
