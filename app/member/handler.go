package main

import (
	"context"
	base "gen/kitex_gen/base"
	member "gen/kitex_gen/member"
	"member/biz/service"
)

// MemberServiceImpl implements the last service interface defined in the IDL.
type MemberServiceImpl struct{}

// CreateMember implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) CreateMember(ctx context.Context, req *member.CreateMemberReq) (resp *member.MemberResp, err error) {
	resp, err = service.NewCreateMemberService(ctx).Run(req)

	return resp, err
}

// DeleteMember implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) DeleteMember(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteMemberService(ctx).Run(req)

	return resp, err
}

// UpdateMember implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) UpdateMember(ctx context.Context, req *member.UpdateMemberReq) (resp *member.MemberResp, err error) {
	resp, err = service.NewUpdateMemberService(ctx).Run(req)

	return resp, err
}

// GetMember implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) GetMember(ctx context.Context, req *base.IdReq) (resp *member.MemberResp, err error) {
	resp, err = service.NewGetMemberService(ctx).Run(req)

	return resp, err
}

// GetMemberList implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) GetMemberList(ctx context.Context, req *member.GetMemberListReq) (resp *member.MemberListResp, err error) {
	resp, err = service.NewGetMemberListService(ctx).Run(req)

	return resp, err
}

// LoginMember implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) LoginMember(ctx context.Context, req *base.CheckAccountReq) (resp *member.MemberResp, err error) {
	resp, err = service.NewLoginMemberService(ctx).Run(req)

	return resp, err
}

// ChangePassword implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) ChangePassword(ctx context.Context, req *member.ChangePasswordReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewChangePasswordService(ctx).Run(req)

	return resp, err
}

// GetMemberIds implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) GetMemberIds(ctx context.Context, req *member.GetMemberListReq) (resp *base.NilResponse, err error) {
	// TODO: Your code here...
	return
}
