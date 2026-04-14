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
	dto := mapper.ToDTO(e)

	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)

	return dto

}
func EntToMessagesSentRecords(e *ent.MessagesSentRecords) *base.MessagesSend {

	if e == nil {
		return nil
	}
	if e.Edges.Messages == nil {
		return nil
	}
	message := e.Edges.Messages

	mapper := utils.NewCopierMapper[base.MessagesSend, ent.MessagesSentRecords]()
	var dto = mapper.ToDTO(e)

	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)

	if e.ReceivedAt != nil {
		dto.ReceivedAt = e.ReceivedAt.Format(time.DateOnly)
	}
	if e.ReadAt != nil {
		dto.ReadAt = e.ReadAt.Format(time.DateOnly)
	}

	dto.Content = *message.Content
	dto.Type = *message.Type
	dto.Title = *message.Title
	dto.FromUserName = *message.FromUserName

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
