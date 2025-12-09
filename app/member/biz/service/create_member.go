package service

import (
	"common/pkg/errno"
	"context"
	Member "gen/kitex_gen/member"
	"member/biz/convert"
	"member/biz/dal/db"
	"member/biz/dal/db/ent/member"
)

type CreateMemberService struct {
	ctx context.Context
} // NewCreateMemberService new CreateMemberService
func NewCreateMemberService(ctx context.Context) *CreateMemberService {
	return &CreateMemberService{ctx: ctx}
}

// Run create note info
func (s *CreateMemberService) Run(req *Member.CreateMemberReq) (resp *Member.MemberResp, err error) {

	ok, _ := db.Client.Member.Query().Where(member.UsernameEQ(req.GetUsername())).Exist(s.ctx)
	if ok {
		return nil, errno.UserAlreadyExistErr
	}
	//ok, _ = db.Client.Member.Query().Where(member.MobileEQ(*req.Membername)).Exist(s.ctx)
	//if ok {
	//	return nil, errno.MemberMobileExistErr
	//}
	save, err := db.Client.Member.Create().
		SetUsername(req.GetUsername()).
		SetPassword(req.GetPassword()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	resp = &Member.MemberResp{
		Data: convert.EntToMember(save),
	}
	return
}
