package service

import (
	"context"
	Base "gen/kitex_gen/base"
	Member "gen/kitex_gen/member"
	"member/biz/dal/db"
	"member/biz/dal/db/ent"
	"member/biz/dal/db/ent/membertag"

	"member/biz/convert"
	"member/biz/dal/db/ent/member"
	"member/biz/dal/db/ent/predicate"
)

type GetMemberIdsService struct {
	ctx context.Context
} // NewGetMemberIdsService new GetMemberIdsService
func NewGetMemberIdsService(ctx context.Context) *GetMemberIdsService {
	return &GetMemberIdsService{ctx: ctx}
}

// Run Update
func (s *GetMemberIdsService) Run(req *Member.GetMemberListReq) (resp *Member.MemberIdsResp, err error) {

	var (
		dataResp []*Base.Member
	)
	var predicates []predicate.Member
	if len(req.GetTags()) > 0 {
		predicates = append(predicates, member.HasMemberTagsWith(membertag.TagIDIn(req.GetTags()...)))
	}

	all, err := db.Client.Member.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(member.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range all {
		profile, err := v.QueryMemberProfile().First(s.ctx)
		if err != nil {
			return nil, err
		}
		dataResp = append(dataResp, convert.EntToMember(v, profile))
	}
	return &Member.MemberIdsResp{
		Data: nil,
	}, nil
}
