package convert

import (
	"common/pkg/utils"
	Base "gen/kitex_gen/base"
	"member/biz/dal/db/ent"
	"time"
)

func EntToMember(e *ent.Member) *Base.Member {

	mapper := utils.NewCopierMapper[Base.Member, ent.Member]()
	var dto = mapper.ToDTO(e)

	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)
	dto.UpdatedAt = e.UpdatedAt.Format(time.DateOnly)
	//dto.LastAt = e.LastAt.Format(time.DateOnly)
	//	Gender:    0,
	//		Birthday:  "",
	return dto

}
