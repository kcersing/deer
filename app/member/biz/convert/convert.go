package convert

import (
	"common/pkg/utils"
	Base "gen/kitex_gen/base"
	"member/biz/dal/db/ent"
	"time"
)

func EntToMember(e *ent.Member, p *ent.MemberProfile) *Base.Member {
	if e == nil || p == nil {
		return nil
	}

	mapper := utils.NewCopierMapper[Base.Member, ent.Member]()
	var dto = mapper.ToDTO(e)

	dto.Intention = p.Intention
	dto.Gender = p.Gender
	dto.Birthday = p.Birthday.Format(time.DateOnly)
	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)
	dto.UpdatedAt = e.UpdatedAt.Format(time.DateOnly)

	//dto.Level = p.Level
	//dto.Source = p.Source
	//dto.LastAt = p.LastAt.Format(time.DateOnly)
	//dto.LastIp = p.LastIp
	//dto.Email = p.Email
	//dto.Wecom = p.Wecom
	//dto.CreatedName = p.CreatedName
	//dto.RelationMid = p.RelationMid
	//dto.RelationMame = p.RelationMame

	return dto

}
