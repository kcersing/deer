package main

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"system/biz/service"
)

// SystemServiceImpl implements the last service interface defined in the IDL.
type SystemServiceImpl struct{}

// CreateApi implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CreateApi(ctx context.Context, req *system.CreateApiReq) (resp *system.ApiResp, err error) {
	resp, err = service.NewCreateApiService(ctx).Run(req)

	return resp, err
}

// UpdateApi implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateApi(ctx context.Context, req *system.UpdateApiReq) (resp *system.ApiResp, err error) {
	resp, err = service.NewUpdateApiService(ctx).Run(req)

	return resp, err
}

// DeleteApi implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteApi(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteApiService(ctx).Run(req)

	return resp, err
}

// ApiList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) ApiList(ctx context.Context, req *system.ApiListReq) (resp *system.ApiListResp, err error) {
	resp, err = service.NewApiListService(ctx).Run(req)

	return resp, err
}

// ApiTree implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) ApiTree(ctx context.Context, req *system.ApiListReq) (resp *system.ApiListResp, err error) {
	resp, err = service.NewApiTreeService(ctx).Run(req)

	return resp, err
}

// CreateMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CreateMenu(ctx context.Context, req *system.CreateMenuReq) (resp *system.MenuResp, err error) {
	resp, err = service.NewCreateMenuService(ctx).Run(req)

	return resp, err
}

// UpdateMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateMenu(ctx context.Context, req *system.UpdateMenuReq) (resp *system.MenuResp, err error) {
	resp, err = service.NewUpdateMenuService(ctx).Run(req)

	return resp, err
}

// DeleteMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteMenu(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteMenuService(ctx).Run(req)

	return resp, err
}

// GetMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) GetMenu(ctx context.Context, req *base.IdReq) (resp *system.MenuResp, err error) {
	resp, err = service.NewGetMenuService(ctx).Run(req)

	return resp, err
}

// MenuList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) MenuList(ctx context.Context, req *system.MenuListReq) (resp *system.MenuListResp, err error) {
	resp, err = service.NewMenuListService(ctx).Run(req)

	return resp, err
}

// MenuTree implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) MenuTree(ctx context.Context, req *system.MenuListReq) (resp *system.MenuListResp, err error) {
	resp, err = service.NewMenuTreeService(ctx).Run(req)

	return resp, err
}

// CreateRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CreateRole(ctx context.Context, req *system.CreateRoleReq) (resp *system.RoleResp, err error) {
	resp, err = service.NewCreateRoleService(ctx).Run(req)

	return resp, err
}

// GetRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) GetRole(ctx context.Context, req *base.IdReq) (resp *system.RoleResp, err error) {
	resp, err = service.NewGetRoleService(ctx).Run(req)

	return resp, err
}

// GetRoleList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) GetRoleList(ctx context.Context, req *system.GetRoleListReq) (resp *system.RoleListResp, err error) {
	resp, err = service.NewGetRoleListService(ctx).Run(req)

	return resp, err
}

// UpdateRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateRole(ctx context.Context, req *system.UpdateRoleReq) (resp *system.RoleResp, err error) {
	resp, err = service.NewUpdateRoleService(ctx).Run(req)

	return resp, err
}

// DeleteRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteRole(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteRoleService(ctx).Run(req)

	return resp, err
}

// CreateRoleMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CreateRoleMenu(ctx context.Context, req *system.CreateMenuAuthReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateRoleMenuService(ctx).Run(req)

	return resp, err
}

// CreateRoleApi implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CreateRoleApi(ctx context.Context, req *system.CreateMenuAuthReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateRoleApiService(ctx).Run(req)

	return resp, err
}

// GetRoleApi implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) GetRoleApi(ctx context.Context, req *base.IdReq) (resp *system.ApiListResp, err error) {
	resp, err = service.NewGetRoleApiService(ctx).Run(req)

	return resp, err
}

// GetRoleMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) GetRoleMenu(ctx context.Context, req *base.IdReq) (resp *system.MenuListResp, err error) {
	resp, err = service.NewGetRoleMenuService(ctx).Run(req)

	return resp, err
}

// CreateDict implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CreateDict(ctx context.Context, req *system.Dict) (resp *system.DictResp, err error) {
	resp, err = service.NewCreateDictService(ctx).Run(req)

	return resp, err
}

// UpdateDict implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateDict(ctx context.Context, req *system.Dict) (resp *system.DictResp, err error) {
	resp, err = service.NewUpdateDictService(ctx).Run(req)

	return resp, err
}

// DeleteDict implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteDict(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteDictService(ctx).Run(req)

	return resp, err
}

// DictList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DictList(ctx context.Context, req *system.DictListReq) (resp *system.DictListResp, err error) {
	resp, err = service.NewDictListService(ctx).Run(req)

	return resp, err
}

// CreateDictht implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CreateDictht(ctx context.Context, req *system.Dictht) (resp *system.DictResp, err error) {
	resp, err = service.NewCreateDicthtService(ctx).Run(req)

	return resp, err
}

// UpdateDictht implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateDictht(ctx context.Context, req *system.Dictht) (resp *system.DictResp, err error) {
	resp, err = service.NewUpdateDicthtService(ctx).Run(req)

	return resp, err
}

// DeleteDictht implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteDictht(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteDicthtService(ctx).Run(req)

	return resp, err
}

// DicthtList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DicthtList(ctx context.Context, req *system.DicthtListReq) (resp *system.DicthtListResp, err error) {
	resp, err = service.NewDicthtListService(ctx).Run(req)

	return resp, err
}

// LogList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) LogList(ctx context.Context, req *system.LogListReq) (resp *system.LogListResp, err error) {
	resp, err = service.NewLogListService(ctx).Run(req)

	return resp, err
}

// DeleteLog implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteLog(ctx context.Context, req *system.DeleteLogReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteLogService(ctx).Run(req)

	return resp, err
}

// VerifyRoleAuth implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) VerifyRoleAuth(ctx context.Context, req *system.VerifyRoleAuthReq) (resp *system.VerifyRoleAuthResp, err error) {
	resp, err = service.NewVerifyRoleAuthService(ctx).Run(req)

	return resp, err
}
