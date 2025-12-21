package convert

import (
	"common/pkg/utils"
	"gen/kitex_gen/base"
	"member/biz/dal/db/ent"
	"time"
)

func EntToMember(e *ent.Member, p *ent.MemberProfile) *base.Member {
	if e == nil || p == nil {
		return nil
	}

	mapper := utils.NewCopierMapper[base.Member, ent.Member]()
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
