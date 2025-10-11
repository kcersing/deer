package main

import (
	"context"
	base "gen/kitex_gen/base"
	member "gen/kitex_gen/member"
)

// MemberServiceImpl implements the last service interface defined in the IDL.
type MemberServiceImpl struct{}

// GetMemberInfo implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) GetMemberInfo(ctx context.Context, req *base.IdReq) (resp *member.MemberResp, err error) {
	// TODO: Your code here...
	return
}
