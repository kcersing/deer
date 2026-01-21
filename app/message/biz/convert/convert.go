package convert

import (
	"common/pkg/utils"
	"gen/kitex_gen/base"
	"message/biz/dal/db/ent"
	"time"
)

func EntToMessages(e *ent.Messages) *base.Messages {
	if e == nil {
		return nil
	}

	mapper := utils.NewCopierMapper[base.Messages, ent.Messages]()
	var dto = mapper.ToDTO(e)

	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)

	return dto

}
