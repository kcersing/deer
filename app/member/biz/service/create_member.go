package service

import (
	"common/pkg/errno"
	"common/pkg/utils"
	"context"
	Member "gen/kitex_gen/member"
	"github.com/pkg/errors"
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

	birthday, err := utils.GetStringDateOnlyZeroTime(req.GetBirthday())
	if err != nil {
		return nil, errno.TimeFormatErr
	}
	ok, _ := db.Client.Member.Query().Where(member.MobileEQ(req.GetMobile())).Exist(s.ctx)
	if ok {
		return nil, errno.UserAlreadyExistErr
	}
	tx, err := db.Client.Tx(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "starting a transaction:")
	}
	save, err := tx.Member.Create().
		SetName(req.GetName()).
		SetUsername(req.GetMobile()).
		SetMobile(req.GetMobile()).
		SetStatus(req.Status).
		SetCreatedID(req.CreatedId).
		SetAvatar(req.GetAvatar()).
		SetCondition(0).
		Save(s.ctx)
	if err != nil {
		return nil, rollback(tx, errors.Wrap(err, "create user failed"))
	}

	profile, err := tx.MemberProfile.Create().
		SetMember(save).
		SetCreatedID(req.CreatedId).
		SetBirthday(birthday).
		SetIntention(req.GetIntention()).
		SetGender(req.GetGender()).
		SetIntention(req.GetIntention()).
		//SetSource(req.GetSource()).
		//SetEmail(req.GetEmail()).
		//SetWecom(req.GetWecom()).
		//SetRelationMid(req.GetRelationMid()).
		//SetRelationMame(req.GetRelationMame()).
		Save(s.ctx)
	if err != nil {
		return nil, rollback(tx, errors.Wrap(err, "create Member Profile failed"))
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	resp = &Member.MemberResp{
		Data: convert.EntToMember(save, profile),
	}
	return
}
