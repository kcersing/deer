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
func EntToMessagesSentRecords(e *ent.MessagesSentRecords) *base.MessagesSend {

	if e == nil {
		return nil
	}

	mapper := utils.NewCopierMapper[base.MessagesSend, ent.MessagesSentRecords]()
	var dto = mapper.ToDTO(e)

	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)

	return dto
}
func EntToSmsSentRecords(e *ent.SmsSentRecords) *base.SmsSend {

	if e == nil {
		return nil
	}

	mapper := utils.NewCopierMapper[base.SmsSend, ent.SmsSentRecords]()
	var dto = mapper.ToDTO(e)

	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)

	return dto
}
