package convert

import (
	"common/pkg/utils"
	"crm/biz/dal/db/ent"
	Base "gen/kitex_gen/base"
	"time"
)

func EntToFollowUpPlan(e *ent.FollowUpPlan) *Base.FollowUpPlan {

	mapper := utils.NewCopierMapper[Base.FollowUpPlan, ent.FollowUpPlan]()
	var dto = mapper.ToDTO(e)

	dto.Time = e.Time.Format(time.DateTime)
	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)
	dto.UpdatedAt = e.UpdatedAt.Format(time.DateOnly)
	return dto
}
func EntToFollowUpRecord(e *ent.FollowUpRecord) *Base.FollowUpRecord {

	mapper := utils.NewCopierMapper[Base.FollowUpRecord, ent.FollowUpRecord]()
	var dto = mapper.ToDTO(e)

	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)
	dto.UpdatedAt = e.UpdatedAt.Format(time.DateOnly)
	return dto
}
func EntToOpportunities(e *ent.Opportunities) *Base.Opportunities {

	mapper := utils.NewCopierMapper[Base.Opportunities, ent.Opportunities]()
	var dto = mapper.ToDTO(e)

	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)
	dto.UpdatedAt = e.UpdatedAt.Format(time.DateOnly)
	return dto
}
