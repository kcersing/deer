package service

import (
	"common/pkg/encrypt"
	"common/pkg/errno"
	"context"
	Base "gen/kitex_gen/base"
	Member "gen/kitex_gen/member"
	"member/biz/dal/db"
	"member/biz/dal/db/ent/member"
)

type LoginMemberService struct {
	ctx context.Context
} // NewLoginMemberService new LoginMemberService
func NewLoginMemberService(ctx context.Context) *LoginMemberService {
	return &LoginMemberService{ctx: ctx}
}

// Run create note info
func (s *LoginMemberService) Run(req *Base.CheckAccountReq) (resp *Member.MemberResp, err error) {
	// Finish your business logic.
	only, err := db.Client.Member.Query().Where(member.UsernameEQ(req.GetUsername())).Only(s.ctx)
	if err != nil {
		return nil, errno.NotFound
	}
	if ok := encrypt.VerifyPassword(req.Password, only.Password); !ok {

		return nil, errno.LoginPasswordErr
	}
	return
}
