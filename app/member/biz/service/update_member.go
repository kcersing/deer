package service

import (
	"common/pkg/errno"
	"context"
	Member "gen/kitex_gen/member"
	"member/biz/convert"
	"member/biz/dal/db"
	"member/biz/dal/db/ent/member"
	"member/biz/dal/db/ent/memberprofile"
	"time"
)

type UpdateMemberService struct {
	ctx context.Context
} // NewUpdateMemberService new UpdateMemberService
func NewUpdateMemberService(ctx context.Context) *UpdateMemberService {
	return &UpdateMemberService{ctx: ctx}
}

// Run Update note info
func (s *UpdateMemberService) Run(req *Member.UpdateMemberReq) (resp *Member.MemberResp, err error) {

	ok, _ := db.Client.Member.Query().Where(member.MobileEQ(req.GetMobile()), member.IDNEQ(req.GetId())).Exist(s.ctx)
	if ok {
		return nil, errno.UserMobileExistErr
	}

	birthday, err := time.Parse(time.DateOnly, req.GetBirthday())
	if err != nil {
		return nil, errno.TimeFormatErr
	}
	only, err := db.Client.Member.UpdateOneID(req.GetId()).
		SetAvatar(req.GetAvatar()).
		SetName(req.GetName()).
		SetStatus(req.GetStatus()).
		//SetGender(req.GetGender()).
		//SetBirthday(birthday).
		//SetDetail(req.GetDetail()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}

	db.Client.MemberProfile.Update().Where(memberprofile.MemberIDEQ(req.GetId())).
		SetGender(req.GetGender()).
		SetBirthday(birthday).
		SetIntention(0).
		//SetEmail(req.GetEmail()).
		//SetWecom(req.GetWecom()).
		//SetSource(req.GetSource()).
		Save(s.ctx)

	resp = &Member.MemberResp{
		Data: convert.EntToMember(only),
	}
	return
}
